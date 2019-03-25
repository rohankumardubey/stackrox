// This is the one place where we're allowed to import directly from 'axios'.
// All other places must use the instance exported here.
// eslint-disable-next-line no-restricted-imports
import axios from 'axios';

export default axios.create({
    timeout: 10000
});

// THE FOLLOWING CODE SNIPPET CAN BE USED TO DEBUG UNIT TESTS,
// IF YOU HAVEN'T MOCKED OUT AXIOS PROPERLY AND ARE GETTING
// CONSOLE ERRORS.
/*
export default {
    get: url => console.log('GET CALLED WITH', url),
    post: (url, data) =>  console.log('POST CALLED WITH', url, data),
    put: (url, data) => console.log('PUT CALLED WITH', url, data),
    patch: (url, data) => console.log('PATCH CALLED WITH', url, data),
    delete: url => console.log('DELETE CALLED WITH', url)
};
*/
