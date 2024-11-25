package errors

import "fmt"

const (
	ValidationSMS string = "2fa_sms"
	ValidationApp string = "2fa_app"
)

const (
	AuthInvalidRequest          string = "invalid_request"
	AuthUnauthorizedClient      string = "unauthorized_client" // app is not allowed to request an authorization code.
	AuthUnsupportedResponseType string = "unsupported_response_type"
	AuthInvalidScope            string = "invalid_scope"
	AuthInvalidGrant            string = "invalid_grant"
	AuthInvalidClient           string = "invalid_client"
	AuthTemporarilyUnavailable  string = "temporarily_unavailable"
	AuthAccessDenied            string = "access_denied"
	AuthNeedValidation          string = "need_validation"
	AuthNeedCaptcha             string = "need_captcha"
	AuthUserDenied              string = "user_denied"
)

type ErrorType int

func (e ErrorType) Error() string {
	return fmt.Sprintf(errorMessagePrefix+"Error with code %d", e)
}

// Doc: https://dev.vk.com/ru/reference/s
const (
	None                                              ErrorType = 0
	UnknownCode                                       ErrorType = 1    // Unknown s occurred
	DisabledCode                                      ErrorType = 2    // App is disabled, enable your app or use test mode (test_mode=1)
	MethodCode                                        ErrorType = 3    // Unknown method passed. http://vk.com/dev/methods
	SignatureCode                                     ErrorType = 4    // Incorrect signature
	AuthCode                                          ErrorType = 5    // Auth s
	TooManyRequestCode                                ErrorType = 6    // Too many requests per second
	PermissionCode                                    ErrorType = 7    // Permission to perform this action is denied
	InvalidRequestCode                                ErrorType = 8    // Invalid query
	FloodCode                                         ErrorType = 9    // Flood control
	InternalServerCode                                ErrorType = 10   // Internal server s
	EnabledInTestCode                                 ErrorType = 11   // In test mode application should be disabled or user should be authorized
	CompileCode                                       ErrorType = 12   // Unable to compile code
	RuntimeCode                                       ErrorType = 13   // Runtime s occurred during code invocation
	CaptchaCode                                       ErrorType = 14   // Captcha needed
	AccessCode                                        ErrorType = 15   // Access denied
	AuthHTTPSCode                                     ErrorType = 16   // HTTP authorization failed, use secure connection
	AuthValidationCode                                ErrorType = 17   // Validation required. https://dev.vk.com/ru/api/validation-required-
	UserDeletedCode                                   ErrorType = 18   // OAuthUser was deleted or banned. https://dev.vk.com/ru/api/validation-required-
	BlockedCode                                       ErrorType = 19   // Content blocked. https://dev.vk.com/ru/api/validation-required-
	MethodPermissionCode                              ErrorType = 20   // Permission to perform this action is denied for non-standalone applications
	MethodAdsCode                                     ErrorType = 21   // Permission to perform this action is allowed only for standalone and OpenAPI applications
	UploadCode                                        ErrorType = 22   // Upload s. Permission to perform this action is allowed only for standalone and OpenAPI applications
	MethodDisabledCode                                ErrorType = 23   // This method was disabled. http://vk.com/dev/methods
	NeedConfirmationCode                              ErrorType = 24   // Confirmation required. https://dev.vk.com/ru/api/confirmation-required-
	NeedTokenConfirmationCode                         ErrorType = 25   // Confirmation required. AccessToken confirmation required. https://dev.vk.com/ru/api/confirmation-required-
	GroupAuthCode                                     ErrorType = 27   // Confirmation required. Group authorization failed. https://dev.vk.com/ru/api/confirmation-required-
	AppAuthCode                                       ErrorType = 28   // Confirmation required. Application authorization failed. https://dev.vk.com/ru/api/confirmation-required-
	RateLimitCode                                     ErrorType = 29   // Rate limit reached. https://dev.vk.com/ru/reference/roadmap
	PrivateProfileCode                                ErrorType = 30   // This profile is private Rate limit reached. https://dev.vk.com/ru/reference/roadmap
	ClientVersionDeprecatedCode                       ErrorType = 34   // TransportClient version deprecated
	ExecutionTimeoutCode                              ErrorType = 36   // Method execution was interrupted due to timeout
	UserBannedCode                                    ErrorType = 37   // OAuthUser was banned
	UnknownApplicationCode                            ErrorType = 38   // Unknown application
	UnknownUserCode                                   ErrorType = 39   // Unknown user
	UnknownGroupCode                                  ErrorType = 40   // Unknown group
	AdditionalSignupRequiredCode                      ErrorType = 41   // Additional signup required
	IPNotAllowedCode                                  ErrorType = 42   // IP is not allowed
	ParamCode                                         ErrorType = 100  // One of the parameters specified was missing or invalid
	ParamAPIIDCode                                    ErrorType = 101  // Invalid application api ID. http://vk.com/apps?act=settings
	LimitsCode                                        ErrorType = 103  // Invalid application api ID. Out of limits
	NotFoundCode                                      ErrorType = 104  // Invalid application api ID. Not found
	SaveFileCode                                      ErrorType = 105  // Invalid application api ID. Couldn't save file
	ActionFailedCode                                  ErrorType = 106  // Invalid application api ID. Unable to process action
	ParamUserIDCode                                   ErrorType = 113  // Invalid user id
	ParamAlbumIDCode                                  ErrorType = 114  // Invalid user id. Invalid album id
	ParamServerCode                                   ErrorType = 118  // Invalid user id. Invalid server
	ParamTitleCode                                    ErrorType = 119  // Invalid user id. Invalid title
	ParamPhotosCode                                   ErrorType = 122  // Invalid user id. Invalid photos
	ParamHashCode                                     ErrorType = 121  // Invalid user id. Invalid hash
	ParamPhotoCode                                    ErrorType = 129  // Invalid user id. Invalid photo
	ParamGroupIDCode                                  ErrorType = 125  // Invalid user id. Invalid group id
	ParamPageIDCode                                   ErrorType = 140  // Invalid user id. Page not found
	AccessPageCode                                    ErrorType = 141  // Invalid user id. Access to page denied
	MobileNotActivatedCode                            ErrorType = 146  // The mobile number of the user is unknown
	InsufficientFundsCode                             ErrorType = 147  // Application has insufficient funds
	AccessMenuCode                                    ErrorType = 148  // Access to the menu of the user denied
	ParamTimestampCode                                ErrorType = 150  // Invalid timestamp
	FriendsListIDCode                                 ErrorType = 171  // Invalid list id
	FriendsListLimitCode                              ErrorType = 173  // Reached the maximum number of lists
	FriendsAddYourselfCode                            ErrorType = 174  // Cannot add user himself as friend
	FriendsAddInEnemyCode                             ErrorType = 175  // Cannot add this user to friends as they have put you on their blacklist
	FriendsAddEnemyCode                               ErrorType = 176  // Cannot add this user to friends as you put him on blacklist
	FriendsAddNotFoundCode                            ErrorType = 177  // Cannot add this user to friends as user not found
	ParamNoteIDCode                                   ErrorType = 180  // Cannot add this user to friends as user not found. Note not found
	AccessNoteCode                                    ErrorType = 181  // Cannot add this user to friends as user not found. Access to note denied
	AccessNoteCommentCode                             ErrorType = 182  // Cannot add this user to friends as user not found. You can't comment this note
	AccessCommentCode                                 ErrorType = 183  // Cannot add this user to friends as user not found. Access to comment denied
	AccessAlbumCode                                   ErrorType = 200  // Access to album denied
	AccessAudioCode                                   ErrorType = 201  // Access to audio denied
	AccessGroupCode                                   ErrorType = 203  // Access to group denied
	AccessVideoCode                                   ErrorType = 204  // Access denied
	AccessMarketCode                                  ErrorType = 205  // Access denied
	WallAccessPostCode                                ErrorType = 210  // Access to wall's post denied
	WallAccessCommentCode                             ErrorType = 211  // Access to wall's comment denied
	WallAccessRepliesCode                             ErrorType = 212  // Access to post comments denied
	WallAccessAddReplyCode                            ErrorType = 213  // Access to status replies denied
	WallAddPostCode                                   ErrorType = 214  // Access to adding post denied
	WallAdsPublishedCode                              ErrorType = 219  // Advertisement post was recently added
	WallTooManyRecipientsCode                         ErrorType = 220  // Too many recipients
	StatusNoAudioCode                                 ErrorType = 221  // OAuthUser disabled track name broadcast
	WallLinksForbiddenCode                            ErrorType = 222  // Hyperlinks are forbidden
	WallReplyOwnerFloodCode                           ErrorType = 223  // Too many replies
	WallAdsPostLimitReachedCode                       ErrorType = 224  // Too many ads posts
	DonutDisabledCode                                 ErrorType = 225  // Donut is disabled
	LikesReactionCanNotBeAppliedCode                  ErrorType = 232  // Reaction can not be applied to the temp
	PollsAccessCode                                   ErrorType = 250  // Access to poll denied
	PollsAnswerIDCode                                 ErrorType = 252  // Invalid answer id
	PollsPollIDCode                                   ErrorType = 251  // Invalid poll id
	PollsAccessWithoutVoteCode                        ErrorType = 253  // Access denied, please vote first
	AccessGroupsCode                                  ErrorType = 260  // Access to the groups list is denied due to the user's privacy settings
	AlbumFullCode                                     ErrorType = 300  // This album is full
	AlbumsLimitCode                                   ErrorType = 302  // This album is full. Albums number limit is reached
	VotesPermissionCode                               ErrorType = 500  // Permission denied. You must enable votes processing in application. http://vk.com/editapp?id={API_ID}&section=payments
	VotesCode                                         ErrorType = 503  // Not enough votes
	NotEnoughMoneyCode                                ErrorType = 504  // Not enough money on owner's balance
	AdsPermissionCode                                 ErrorType = 600  // Permission denied. You have no access to operations specified with given temp
	WeightedFloodCode                                 ErrorType = 601  // Permission denied. You have requested too many action this day. Try later
	AdsPartialSuccessCode                             ErrorType = 602  // Some part of the query has not been completed
	AdsSpecificCode                                   ErrorType = 603  // Some ads s occurred
	AdsDomainInvalidCode                              ErrorType = 604  // Invalid domain
	AdsDomainForbiddenCode                            ErrorType = 605  // Domain is forbidden
	AdsDomainReservedCode                             ErrorType = 606  // Domain is reserved
	AdsDomainOccupiedCode                             ErrorType = 607  // Domain is occupied
	AdsDomainActiveCode                               ErrorType = 608  // Domain is active
	AdsDomainAppInvalidCode                           ErrorType = 609  // Domain app is invalid
	AdsDomainAppForbiddenCode                         ErrorType = 610  // Domain app is forbidden
	AdsApplicationMustBeVerifiedCode                  ErrorType = 611  // Application must be verified
	AdsApplicationMustBeInDomainsListCode             ErrorType = 612  // Application must be in domains list of site of ad unit
	AdsApplicationBlockedCode                         ErrorType = 613  // Application is blocked
	AdsDomainTypeForbiddenInCurrentOfficeCode         ErrorType = 614  // Domain of type specified is forbidden in current office type
	AdsDomainGroupInvalidCode                         ErrorType = 615  // Domain group is invalid
	AdsDomainGroupForbiddenCode                       ErrorType = 616  // Domain group is forbidden
	AdsDomainAppBlockedCode                           ErrorType = 617  // Domain app is blocked AdsDomainGroupNotOpen ErrorType = 618 	// Domain group is not open
	AdsDomainGroupNotPossibleJoinedCode               ErrorType = 619  // Domain group is not possible to be joined to adsweb
	AdsDomainGroupBlockedCode                         ErrorType = 620  // Domain group is blocked
	AdsDomainGroupLinksForbiddenCode                  ErrorType = 621  // Domain group has restriction: links are forbidden
	AdsDomainGroupExcludedFromSearchCode              ErrorType = 622  // Domain group has restriction: excluded from search
	AdsDomainGroupCoverForbiddenCode                  ErrorType = 623  // Domain group has restriction: cover is forbidden
	AdsDomainGroupWrongCategoryCode                   ErrorType = 624  // Domain group has wrong category
	AdsDomainGroupWrongNameCode                       ErrorType = 625  // Domain group has wrong name
	AdsDomainGroupLowPostsReachCode                   ErrorType = 626  // Domain group has low posts reach
	AdsDomainGroupWrongClassCode                      ErrorType = 627  // Domain group has wrong class
	AdsDomainGroupCreatedRecentlyCode                 ErrorType = 628  // Domain group is created recently
	AdsObjectDeletedCode                              ErrorType = 629  // Object deleted
	AdsLookalikeRequestAlreadyInProgressCode          ErrorType = 630  // Lookalike query with same source already in progress
	AdsLookalikeRequestsLimitCode                     ErrorType = 631  // Max count of lookalike requests per day reached
	AdsAudienceTooSmallCode                           ErrorType = 632  // Given audience is too small
	AdsAudienceTooLargeCode                           ErrorType = 633  // Given audience is too large
	AdsLookalikeAudienceSaveAlreadyInProgressCode     ErrorType = 634  // Lookalike query audience save already in progress
	AdsLookalikeSavesLimitCode                        ErrorType = 635  // Max count of lookalike query audience saves per day reached
	AdsRetargetingGroupsLimitCode                     ErrorType = 636  // Max count of retargeting groups reached
	AdsDomainGroupActiveNemesisPunishmentCode         ErrorType = 637  // Domain group has active nemesis punishment
	GroupChangeCreatorCode                            ErrorType = 700  // Cannot edit creator role
	GroupNotInClubCode                                ErrorType = 701  // OAuthUser should be in club
	GroupTooManyOfficersCode                          ErrorType = 702  // Too many officers in club
	GroupNeed2faCode                                  ErrorType = 703  // You need to enable 2FA for this action
	GroupHostNeed2faCode                              ErrorType = 704  // OAuthUser needs to enable 2FA for this action
	GroupTooManyAddressesCode                         ErrorType = 706  // Too many addresses in club
	GroupAppIsNotInstalledInCommunityCode             ErrorType = 711  // Application is not installed in community
	GroupInvalidInviteLinkCode                        ErrorType = 714  // Invite link is invalid - expired, deleted or not exists
	VideoAlreadyAddedCode                             ErrorType = 800  // This video is already added
	VideoCommentsClosedCode                           ErrorType = 801  // Comments for this video are closed
	MessagesUserBlockedCode                           ErrorType = 900  // Can't send messages for users from blacklist
	MessagesDenySendCode                              ErrorType = 901  // Can't send messages for users without permission
	MessagesPrivacyCode                               ErrorType = 902  // Can't send messages to this user due to their privacy settings
	MessagesTooOldPtsCode                             ErrorType = 907  // Value of ts or pts is too old
	MessagesTooNewPtsCode                             ErrorType = 908  // Value of ts or pts is too new
	MessagesEditExpiredCode                           ErrorType = 909  // Can't edit this message, because it's too old
	MessagesTooBigCode                                ErrorType = 910  // Can't sent this message, because it's too big
	MessagesKeyboardInvalidCode                       ErrorType = 911  // Keyboard format is invalid
	MessagesChatBotFeatureCode                        ErrorType = 912  // This is a chat bot feature, change this status in settings
	MessagesTooLongForwardsCode                       ErrorType = 913  // Too many forwarded messages
	MessagesTooLongMessageCode                        ErrorType = 914  // EventType is too long
	MessagesChatUserNoAccessCode                      ErrorType = 917  // You don't have access to this chat
	MessagesCantSeeInviteLinkCode                     ErrorType = 919  // You can't see invite link for this chat
	MessagesEditKindDisallowedCode                    ErrorType = 920  // Can't edit this kind of message
	MessagesCantFwdCode                               ErrorType = 921  // Can't forward these messages
	MessagesCantDeleteForAllCode                      ErrorType = 924  // Can't delete this message for everybody
	MessagesChatNotAdminCode                          ErrorType = 925  // You are not admin of this chat
	MessagesChatNotExistCode                          ErrorType = 927  // Chat does not exist
	MessagesCantChangeInviteLinkCode                  ErrorType = 931  // You can't change invite link for this chat
	MessagesGroupPeerAccessCode                       ErrorType = 932  // Your community can't interact with this peer
	MessagesChatUserNotInChatCode                     ErrorType = 935  // OAuthUser not found in chat
	MessagesContactNotFoundCode                       ErrorType = 936  // Contact not found
	MessagesMessageRequestAlreadySendCode             ErrorType = 939  // EventType query already send
	MessagesTooManyPostsCode                          ErrorType = 940  // Too many posts in messages
	MessagesCantPinOneTimeStoryCode                   ErrorType = 942  // Cannot pin one-time story
	MessagesCantUseIntentCode                         ErrorType = 943  // Cannot use this intent
	MessagesLimitIntentCode                           ErrorType = 944  // Limits overflow for this intent
	MessagesChatDisabledCode                          ErrorType = 945  // Chat was disabled
	MessagesChatNotSupportedCode                      ErrorType = 946  // Chat not support
	MessagesMemberAccessToGroupDeniedCode             ErrorType = 947  // Can't add user to chat, because user has no access to group
	MessagesEditPinnedCode                            ErrorType = 949  // Can't edit pinned message yet
	MessagesReplyTimedOutCode                         ErrorType = 950  // Can't send message, reply timed out
	MessagesAccessDonutChatCode                       ErrorType = 962  // You can't access donut chat without subscription
	MessagesAccessWorkChatCode                        ErrorType = 967  // This user can't be added to the work chat, as they aren't an employe
	MessagesCantForwardedCode                         ErrorType = 969  // EventType cannot be forwarded
	MessagesPinExpiringMessageCode                    ErrorType = 970  // Cannot pin an expiring message
	ParamPhoneCode                                    ErrorType = 1000 // Invalid phone number
	PhoneAlreadyUsedCode                              ErrorType = 1004 // This phone number is used by another user
	AuthFloodCode                                     ErrorType = 1105 // Too many auth attempts, try again later
	AuthDelayCode                                     ErrorType = 1112 // Processing. Try later
	AnonymousTokenExpiredCode                         ErrorType = 1114 // Anonymous token has expired
	AnonymousTokenInvalidCode                         ErrorType = 1116 // Anonymous token is invalid
	AuthAccessTokenHasExpiredCode                     ErrorType = 1117 // Access token has expired
	AuthAnonymousTokenIPMismatchCode                  ErrorType = 1118 // Anonymous token ip mismatch
	ParamDocIDCode                                    ErrorType = 1150 // Invalid document id
	ParamDocDeleteAccessCode                          ErrorType = 1151 // Access to document deleting is denied
	ParamDocTitleCode                                 ErrorType = 1152 // Invalid document title
	ParamDocAccessCode                                ErrorType = 1153 // Access to document is denied
	PhotoChangedCode                                  ErrorType = 1160 // Original photo was changed
	TooManyListsCode                                  ErrorType = 1170 // Too many feed lists
	AppsAlreadyUnlockedCode                           ErrorType = 1251 // This achievement is already unlocked
	AppsSubscriptionNotFoundCode                      ErrorType = 1256 // Subscription not found
	AppsSubscriptionInvalidStatusCode                 ErrorType = 1257 // Subscription is in invalid status
	InvalidAddressCode                                ErrorType = 1260 // Invalid screen name
	CommunitiesCatalogDisabledCode                    ErrorType = 1310 // Catalog is not available for this user
	CommunitiesCategoriesDisabledCode                 ErrorType = 1311 // Catalog categories are not available for this user
	MarketRestoreTooLateCode                          ErrorType = 1400 // Too late for restore
	MarketCommentsClosedCode                          ErrorType = 1401 // Comments for this market are closed
	MarketAlbumNotFoundCode                           ErrorType = 1402 // Album not found
	MarketItemNotFoundCode                            ErrorType = 1403 // Item not found
	MarketItemAlreadyAddedCode                        ErrorType = 1404 // Item already added to album
	MarketTooManyItemsCode                            ErrorType = 1405 // Too many items
	MarketTooManyItemsInAlbumCode                     ErrorType = 1406 // Too many items in album
	MarketTooManyAlbumsCode                           ErrorType = 1407 // Too many albums
	MarketItemHasBadLinksCode                         ErrorType = 1408 // Item has bad links in description
	MarketShopNotEnabledCode                          ErrorType = 1409 // Extended market not enabled
	MarketGroupingItemsWithDifferentPropertiesCode    ErrorType = 1412 // Grouping items with different properties
	MarketGroupingAlreadyHasSuchVariantCode           ErrorType = 1413 // Grouping already has such variant
	MarketVariantNotFoundCode                         ErrorType = 1416 // Variant not found
	MarketPropertyNotFoundCode                        ErrorType = 1417 // Property not found
	MarketGroupingMustContainMoreThanOneItemCode      ErrorType = 1425 // Grouping must have two or more items
	MarketGroupingItemsMustHaveDistinctPropertiesCode ErrorType = 1426 // Item must have distinct properties
	MarketOrdersNoCartItemsCode                       ErrorType = 1427 // Cart is empty
	MarketInvalidDimensionsCode                       ErrorType = 1429 // Specify width, length, height and weight all together
	MarketCantChangeVkpayStatusCode                   ErrorType = 1430 // VK Pay status can not be changed
	MarketShopAlreadyEnabledCode                      ErrorType = 1431 // Market was already enabled in this group
	MarketShopAlreadyDisabledCode                     ErrorType = 1432 // Market was already disabled in this group
	MarketPhotosCropInvalidFormatCode                 ErrorType = 1433 // Invalid image crop format
	MarketPhotosCropOverflowCode                      ErrorType = 1434 // Crop bottom right corner is outside of the image
	MarketPhotosCropSizeTooLowCode                    ErrorType = 1435 // Crop size is less than the minimum
	MarketNotEnabledCode                              ErrorType = 1438 // Market not enabled
	MarketCartEmptyCode                               ErrorType = 1427 // Cart is empty
	MarketSpecifyDimensionsCode                       ErrorType = 1429 // Specify width, length, height and weight all together
	VKPayStatusCode                                   ErrorType = 1430 // VK Pay status can not be changed
	MarketAlreadyEnabledCode                          ErrorType = 1431 // Market was already enabled in this group
	MarketAlreadyDisabledCode                         ErrorType = 1432 // Market was already disabled in this group
	MainAlbumCantHiddenCode                           ErrorType = 1446 // Main album can not be hidden
	StoryExpiredCode                                  ErrorType = 1600 // Story has already expired
	StoryIncorrectReplyPrivacyCode                    ErrorType = 1602 // Incorrect reply privacy
	PrettyCardsCardNotFoundCode                       ErrorType = 1900 // Card not found
	PrettyCardsTooManyCardsCode                       ErrorType = 1901 // Too many cards
	PrettyCardsCardIsConnectedToPostCode              ErrorType = 1902 // Card is connected to post
	CallbackServersLimitCode                          ErrorType = 2000 // Servers number limit is reached
	StickersNotPurchasedCode                          ErrorType = 2100 // Stickers are not purchased
	StickersTooManyFavoritesCode                      ErrorType = 2101 // Too many favorite stickers
	StickersNotFavoriteCode                           ErrorType = 2102 // Stickers are not favorite
	WallCheckLinkCantDetermineSourceCode              ErrorType = 3102 // Specified link is incorrect (can't find source)
	RecaptchaCode                                     ErrorType = 3300 // Recaptcha needed
	PhoneValidationCode                               ErrorType = 3301 // Phone validation needed
	PasswordValidationCode                            ErrorType = 3302 // Password validation needed
	OtpAppValidationCode                              ErrorType = 3303 // Otp app validation needed
	EmailConfirmationCode                             ErrorType = 3304 // Email confirmation needed
	AssertVotesCode                                   ErrorType = 3305 // Assert votes
	TokenExtensionCode                                ErrorType = 3609 // AccessToken extension required
	UserDeactivatedCode                               ErrorType = 3610 // OAuthUser is deactivated
	ServiceDeactivatedCode                            ErrorType = 3611 // Service is deactivated for user
	AliExpressTagCode                                 ErrorType = 3800 // Can't set AliExpress tag to this code of temp
	InvalidUploadResponseCode                         ErrorType = 5701 // Invalid upload response
	InvalidUploadHashCode                             ErrorType = 5702 // Invalid upload hash
	InvalidUploadUserCode                             ErrorType = 5703 // Invalid upload user
	InvalidUploadGroupCode                            ErrorType = 5704 // Invalid upload group
	InvalidCropDataCode                               ErrorType = 5705 // Invalid crop data
	ToSmallAvatarCode                                 ErrorType = 5706 // To small avatar
	PhotoNotFoundCode                                 ErrorType = 5708 // Photo not found
	InvalidPhotoCode                                  ErrorType = 5709 // Invalid Photo
	InvalidHashCode                                   ErrorType = 5710 // Invalid hash
)

// PaymentsErrorCode
//
//	Errors 100-999 are specified by the app. Such errors always include an error description.
//	See https://dev.vk.com/ru/api/payments/notifications/overview
type PaymentsErrorCode int

const (
	PaymentsCommonError            PaymentsErrorCode = 1  // Critical: true/false.
	PaymentsTemporaryDatabaseError PaymentsErrorCode = 2  // Critical: false.
	PaymentsBadSignatures          PaymentsErrorCode = 10 // Critical: true.
	// BadRequest Request parameters do not comply with the specification.
	// No required fields in the request.
	// Other request integrity errors.
	PaymentsBadRequest        PaymentsErrorCode = 11 // Critical: true.
	PaymentsProductNotExist   PaymentsErrorCode = 20 // Critical: true.
	PaymentsProductOutOfStock PaymentsErrorCode = 21 // Critical: true.
	PaymentsUserNotExist      PaymentsErrorCode = 22 // Critical: true.
)
