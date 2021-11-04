package pgsearch

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/parse"
	pkgSearch "github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/enumregistry"
	"github.com/stackrox/rox/pkg/stringutils"
	//"github.com/stackrox/rox/pkg/search/enumregistry"
	"github.com/stackrox/rox/pkg/utils"
)

type queryFunction func(table string, field *pkgSearch.Field, value string, queryModifiers ...pkgSearch.QueryModifier) (*QueryEntry, error)

var datatypeToQueryFunc = map[v1.SearchDataType]queryFunction{
	v1.SearchDataType_SEARCH_STRING:   newStringQuery,
	v1.SearchDataType_SEARCH_BOOL:     newBoolQuery,
	v1.SearchDataType_SEARCH_NUMERIC:  newNumericQuery,
	v1.SearchDataType_SEARCH_DATETIME: newTimeQuery,
	v1.SearchDataType_SEARCH_ENUM:     newEnumQuery,
	// Map type is handled specially.
}

func matchFieldQuery(table string, field *pkgSearch.Field, value string) (*QueryEntry, error) {
	// Special case: wildcard
	if stringutils.MatchesAny(value, pkgSearch.WildcardString, pkgSearch.NullString) {
		return handleExistenceQueries(table, field, value), nil
	}

	trimmedValue, modifiers := pkgSearch.GetValueAndModifiersFromString(value)
	return datatypeToQueryFunc[field.GetType()](table, field, trimmedValue, modifiers...)
}

func handleExistenceQueries(table string, field *pkgSearch.Field, value string) *QueryEntry {
	lastElem := field.LastElem()
	elemPath := GenerateShortestElemPath(table, field.Elems)

	root := field.TopLevelValue()
	switch value {
	case pkgSearch.WildcardString:
		if root != "" {
			return &QueryEntry{
				Query: fmt.Sprintf("%s is not null", root),
			}
		} else {
			return &QueryEntry{
				Query: fmt.Sprintf("%s ? $$", elemPath),
				Values: []interface{}{lastElem.ProtoJSONName},
			}
		}
	case pkgSearch.NullString:
		if root != "" {
			return &QueryEntry{
				Query: fmt.Sprintf("%s is null", root),
			}
		} else {
			return &QueryEntry{
				Query: fmt.Sprintf("not %s ? $$", elemPath),
				Values: []interface{}{lastElem.ProtoJSONName},
			}
		}
	default:
		log.Fatalf("existence query for value %s is not currently handled", value)
	}
	return nil
}

//func getWildcardQuery(field string) *query.WildcardQuery {
//	wq := bleve.NewWildcardQuery("*")
//	wq.SetField(field)
//	return wq
//}

func RenderFinalPath(elemPath string, field string) string {
	if elemPath == "" {
		return field + " "
	}
	return fmt.Sprintf("%s ->>'%s' ", elemPath, field)
}

func newStringQuery(table string, field *pkgSearch.Field, value string, queryModifiers ...pkgSearch.QueryModifier) (*QueryEntry, error) {
	if len(value) == 0 {
		return nil, errors.New("value in search query cannot be empty")
	}

	lastElem := field.LastElem()
	elemPath := GenerateShortestElemPath(table, field.Elems)

	root := field.TopLevelValue()
	if root == "" {
		root = RenderFinalPath(elemPath, lastElem.ProtoJSONName)
	}

	if len(queryModifiers) == 0 {
		return &QueryEntry{
			Query:  root + " ilike $$",
			Values: []interface{}{value + "%"},
		}, nil
	}
	if queryModifiers[0] == pkgSearch.AtLeastOne {
		panic("I dont think this is used")
	}
	var negationString string
	if negated := queryModifiers[0] == pkgSearch.Negation; negated {
		negationString = "!"
		queryModifiers = queryModifiers[1:]
	}

	switch queryModifiers[0] {
	case pkgSearch.Regex:
		return &QueryEntry{
			Query:  root + fmt.Sprintf(" %s~* $$", negationString),
			Values: []interface{}{value},
		}, nil
	case pkgSearch.Equality:
		return &QueryEntry{
			Query:  root + fmt.Sprintf(" %s= $$", negationString),
			Values: []interface{}{value},
		}, nil
	}
	err := errors.Errorf("unknown query modifier: %s", queryModifiers[0])
	utils.Should(err)
	return nil, err
}

func parseLabel(label string) (string, string) {
	spl := strings.SplitN(label, "=", 2)
	if len(spl) < 2 {
		return spl[0], ""
	}
	return spl[0], spl[1]
}

func newBoolQuery(table string, field *pkgSearch.Field, value string, modifiers ...pkgSearch.QueryModifier) (*QueryEntry, error) {
	if len(modifiers) > 0 {
		return nil, errors.Errorf("modifiers for bool query not allowed: %+v", modifiers)
	}
	_, err := parse.FriendlyParseBool(value)
	if err != nil {
		return nil, err
	}
	// explicitly apply equality check
	return newStringQuery(table, field, value, pkgSearch.Equality)
}

func newEnumQuery(table string, field *pkgSearch.Field, value string, queryModifiers ...pkgSearch.QueryModifier) (*QueryEntry, error) {
	var enumValues []int32
	if len(queryModifiers) > 2 {
		return nil, errors.Errorf("unsupported: more than 2 query modifiers for enum query: %+v", queryModifiers)
	}
	var equality bool
	switch len(queryModifiers) {
	case 2:
		if queryModifiers[0] == pkgSearch.Negation && queryModifiers[1] == pkgSearch.Regex {
			re, err := regexp.Compile(value)
			if err != nil {
				return nil, errors.Wrap(err, "invalid regex")
			}

			enumValues = enumregistry.GetComplementOfValuesMatchingRegex(field.FieldPath, re)
			break
		}
		if queryModifiers[0] == pkgSearch.Negation && queryModifiers[1] == pkgSearch.Equality {
			enumValues = enumregistry.GetComplementByExactMatches(field.FieldPath, value)
			break
		}
		return nil, errors.Errorf("unsupported: invalid combination of query modifiers for enum query: %+v", queryModifiers)
	case 1:
		switch queryModifiers[0] {
		case pkgSearch.Negation:
			enumValues = enumregistry.GetComplement(field.FieldPath, value)
		case pkgSearch.Regex:
			re, err := regexp.Compile(value)
			if err != nil {
				return nil, errors.Wrap(err, "invalid regex")
			}
			enumValues = enumregistry.GetValuesMatchingRegex(field.FieldPath, re)
		case pkgSearch.Equality:
			enumValues = enumregistry.GetExactMatches(field.FieldPath, value)
			equality = true
		default:
			return nil, errors.Errorf("unsupported query modifier for enum query: %v", queryModifiers[0])
		}
	case 0:
		prefix, value := parseNumericPrefix(value)
		if prefix == "" {
			equality = true
		}
		enumValues = enumregistry.Get(field.FieldPath, value)
		if len(enumValues) == 0 {
			return NewFalseQuery(), nil
		}

		// Equality means no numeric cast required, and could benefit from hash indexes
		if equality {
			var queries []string
			var values []interface{}
			for _, s := range enumValues {
				entry, err := newStringQuery(table, field, strconv.Itoa(int(s)), pkgSearch.Equality)
				if err != nil {
					return nil, err
				}
				queries = append(queries, entry.Query)
				values = append(values, entry.Values...)
			}
			return &QueryEntry{
				Query:  fmt.Sprintf("(%s)", strings.Join(queries, " or ")),
				Values: values,
			}, nil
		}

		var queries []string
		var values []interface{}
		for _, s := range enumValues {
			entry := createNumericQuery(table, field, prefix, floatPtr(float64(s)))
			queries = append(queries, entry.Query)
			values = append(values, entry.Values...)
		}
		return &QueryEntry{
			Query:  fmt.Sprintf("(%s)", strings.Join(queries, " or ")),
			Values: values,
		}, nil
	}

	if len(enumValues) == 0 {
		return nil, fmt.Errorf("could not find corresponding enum at field %q with value %q and modifiers %+v", field, value, queryModifiers)
	}

	var queries []string
	var values []interface{}
	for _, s := range enumValues {
		entry := createNumericQuery(table, field, "=", floatPtr(float64(s)))
		queries = append(queries, entry.Query)
		values = append(values, entry.Values...)
	}
	return &QueryEntry{
		Query:  fmt.Sprintf("(%s)", strings.Join(queries, " or ")),
		Values: values,
	}, nil
}