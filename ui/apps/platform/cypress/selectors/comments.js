import scopeSelectors from '../helpers/scopeSelectors';

const editSelectors = {
    textArea: 'textarea[data-testid="comment-textarea"]',
    saveButton: 'button[data-testid="save-comment-button"]',
    error: '.pf-c-form__helper-text',
    cancelButton: '[data-testid="cancel-comment-editing-button"]',
};

const commentSelectors = {
    editButton: 'button[data-testid="edit-comment-button"]',
    deleteButton: 'button[data-testid="delete-comment-button"]',
    userName: '[data-testid="comment-header-title"] ',
    dateAndEditedStatus: '[data-testid="comment-header-subtitle"]',
    message: '[data-testid="comment-message"]',
    links: 'a[data-testid="comment-link"]',
    ...editSelectors, // when editing existing comment it's applicable
};

const selectors = {
    allComments: '[data-testid="comment"]',
    lastComment: scopeSelectors('[data-testid="comment"]:last', commentSelectors),

    newButton: 'button[data-testid="new-comment-button"]',

    newComment: scopeSelectors('[data-testid="new-comment"]', editSelectors),
};

export const violationCommentsSelectors = scopeSelectors(
    '[data-testid="violation-comments"]',
    selectors
);

export const processCommentsSelectors = scopeSelectors(
    '[data-testid="process-comments"]',
    selectors
);

export const commentsDialogSelectors = scopeSelectors('.pf-c-modal-box', {
    deleteButton: 'button:contains("Delete")',
});
