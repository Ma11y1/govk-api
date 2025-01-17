package objects

type PagesWikipage struct {
	CreatorID   int    `json:"creator_id"`   // Page creator ID
	CreatorName int    `json:"creator_name"` // Page creator name
	EditorID    int    `json:"editor_id"`    // Last editor ID
	EditorName  string `json:"editor_name"`  // Last editor name
	GroupID     int    `json:"group_id"`     // Community ID
	ID          int    `json:"id"`           // Page ID
	Title       string `json:"title"`        // Page title
	Views       int    `json:"views"`        // Views number
	WhoCanEdit  int    `json:"who_can_edit"` // Edit settings of the page
	WhoCanView  int    `json:"who_can_view"` // View settings of the page
}

type PagesWikipageFull struct {
	// Date when the page has been created in Unixtime.
	Created int `json:"created"`

	// Page creator ID.
	CreatorID int `json:"creator_id"`

	// Information whether current user can edit the page.
	CurrentUserCanEdit BoolInt `json:"current_user_can_edit"`

	// Information whether current user can edit the page access settings.
	CurrentUserCanEditAccess BoolInt `json:"current_user_can_edit_access"`

	// Date when the page has been edited in Unixtime.
	Edited int `json:"edited"`

	// Last editor ID.
	EditorID int `json:"editor_id"`

	// Page ID.
	PageID int `json:"page_id"`

	// Community ID.
	GroupID int `json:"group_id"`

	// Page content, HTML.
	HTML string `json:"html"`

	// Page ID.
	ID int `json:"id"`

	// Page content, wiki.
	Source string `json:"source"`

	// Page title.
	Title string `json:"title"`

	// url of the page preview.
	ViewURL string `json:"view_url"`

	// Views number.
	Views int `json:"views"`

	// Edit settings of the page.
	WhoCanEdit int `json:"who_can_edit"`

	// View settings of the page.
	WhoCanView     int `json:"who_can_view"`
	VersionCreated int `json:"version_created"`
}

type PagesWikipageHistory struct {
	ID         int    `json:"id"`
	Length     int    `json:"length"`
	Edited     int    `json:"edited"`
	EditorID   int    `json:"editor_id"`
	EditorName string `json:"editor_name"`
}
