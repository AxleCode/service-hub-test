package constants

const (
	StaticOTP = "123456"

	DefaultIsActiveValue = true

	DefaultBaseDecimal = 10
	DefaultBitSize     = 64

	DefaultAllowHeaderToken         = "token"
	DefaultAllowHeaderRefreshToken  = "refresh-token"
	DefaultAllowHeaderAuthorization = "Authorization"
	DefaultIsAllLanguage            = "is-all-language"
	DefaultLanguage                 = "language"

	DefaultOrderValue = "created_at DESC"

	LanguageID = "id"
	LanguageEN = "en"

	RoleAdministator = 1

	QuestionTypeScale = "scale"
	QuestionTypeText  = "text"

	PointCredit = "credit"
	PointDebit  = "debit"

	OnSchedule  = "1"
	OffSchedule = "0"

	CreatedAppointmentSuccess = "1"
	CreatedAppointmentFailed  = "0"
	CancelAppointmentSuccess  = "1"
	CancelAppointmentFailed   = "0"

	// status on db.
	AppointmentReservation     = "1"
	AppointmentStatusScheduled = "scheduled"
	AppointmentStatusDone      = "done"
	AppointmentStatusCanceled  = "canceled"
	AppointmentStatusCheckIn   = "check in"
	// status on simrs
	// status : 0 = masih reservasi, 1 = sedang dijalankan (sudah check in), 2 = selesai, 3 = batal.
	SIMRSAppointmentStatusReserved = 0
	SIMRSAppointmentStatusCheckIn  = 1
	SIMRSAppointmentStatusDone     = 2
	SIMRSAppointmentStatusCanceled = 3

	PatientExist           = "1"
	DoctorScheduleUnActive = "0"
	DoctorScheduleActive   = "1"

	MenuStatusActive   = "active"
	MenuStatusUnactive = "unactive"

	GroupCorporate   = "Mora Group"
	GroupCorporateID = "8bfd0df8-f035-4329-af6b-fde686ad1f43"
	SuperadminID     = "81322b1d-8211-429f-9e0b-8ac6e1315043"
	AdminID          = "6afd4ed5-03e5-4015-9274-6f5061a17733"

	MoraCommissionRate = 8.00
	PphCommisionTax    = 15.00

	PowerproHealthCheckPath = "/api/1.0/json/Hello?token="

	IdentifierTypeEmail          = "email"
	IdentifierTypePhone          = "phone"
	IdentifierTypeEmployeeNumber = "employee-number"
)

const (
	DeviceIOS             = "IOS"
	DeviceAndroid         = "Android"
	DeviceWeb             = "web"
	FirebaseStatusSuccess = "success"
	FirebaseStatusFailed  = "failed"

	NotificationAppointment   = "4c3ebfc1-a360-4929-87f3-4da6d817ca0e"
	NotificationBroadcast     = "81dcb6c1-94eb-4975-b208-7376b9f10d3e"
	NotificationPayment       = "f18e6ef0-c334-4ea2-baff-75299f91c4b1"
	NotificationSurvey        = "e0eab3b9-7c0a-4910-9bd9-501d73cacf2c"
	NotificationRegristration = "390279f5-3e12-4e1f-99a5-02e535dc8fab"
	SurveyCategoryAppointment = "ea7a4623-5952-4630-8664-983bad7e69dd"

	NotificationStatusCreated = "created"
	NotificationStatusSent    = "sent"

	NotificationCategoryAppointment   = "appointment"
	NotificationCategoryBroadcast     = "broadcast"
	NotificationCategoryPayment       = "payment"
	NotificationCategorySurvey        = "survey"
	NotificationCategoryRegristration = "registration"

	TrxRegisPatient      = "registration-patient"
	PaymentStatusPaid    = "paid"
	PaymentStatusUnpaid  = "unpaid"
	PaymentStatusFailed  = "failed"
	PaymentStatusExpired = "expired"

	PaymentMethodVA      = "06f243bd-3f1f-4960-9d0c-b97c3d9ff870" // Transfer Virtual
	PaymentMethodCC      = "8cf0c909-46a2-4c66-a63a-543b17b8f690" // Credit Card
	PaymentMethodEWallet = "c183d395-a1d5-4772-8f7a-7ea61a7521fa" // Electronic Wallet
	PaymentMethodDebit   = "0e08ccc2-a6a1-491f-9b0e-1a66366bc0f7" // Direct Debit
	PaymentMethodRetail  = "f2ffeada-f4fd-4ae0-b2d5-37887d1cd8f7" // Retail

	GenderLakiLaki  = "Laki-laki"
	GenderPerempuan = "Perempuan"
)

const (
	XenditEWalletStatusSuccess = "SUCCEEDED"
	XenditEWalletStatusFailed  = "FAILED"

	XenditDebitStatusSuccess = "COMPLETED"
	XenditDebitStatusFailed  = "FAILED"
	XenditEventDebitExpire   = "payment_method.expiry.expired"
	XenditEventDebitExpiring = "payment_method.expiry.expiring"
	XenditEventDebitPayment  = "direct_debit.payment"

	XenditDebitCardEXPIRED = "EXPIRED"
)

type ChannelCodeEnum string

type FailureCodeEwallet string

const (
	// --- Status ENUM ----.
	StatusActive      = "active"
	StatusDeleted     = "deleted"
	StatusComplete    = "complete"
	StatusClosed      = "closed"
	StatusUnqualified = "Unqualified"

	StatusSuspend = "suspend"

	StatusDraft            = "draft"
	StatusScheduledPublish = "scheduled publish"
	StatusPublished        = "published"

	StatusRejected           = "reject"
	StatusPending            = "pending"
	StatusWaitingForApproval = "waiting for approval"
	StatusApproved           = "approved"

	StatusTerminated = "terminated"

	// --- Created By Constants Value ---.
	CreatedByTemporaryBySystem = "temporary by system"

	// --- Delimiter ---.
	DefaultDelimiterStringValue = "|"

	// --- Boolean ---.
	DefaultExpiredStatus  = false
	DefaultDoneStatus     = false
	DefaultArchivedStatus = false
	DefaultSeenStatus     = false

	// --- Board Task Priority ---
	BoardTaskPriorityLow    = "Low"
	BoardTaskPriorityMedium = "Medium"
	BoardTaskPriorityHigh   = "High"
	BoardTaskNoPriority     = "No Priority"

	// --- Board Task Filter ---
	NoAssignee    = "No Assignee"
	NoLabel       = "No Label"
	TaskNoDueDate = "No Dates"
	TaskOverDue   = "Overdue"
	TaskNextDay   = "Next Day"
	TaskNextWeek  = "Next Week"
	TaskNextMonth = "Next Month"

	// --- Board Template ---
	BoardTemplateKanban = "811ab326-8e05-409c-af7f-469255fa72f4"
)

// Time Format.
const (
	TimeDateFormat             = "2006-01-02"
	TimeWithSecondFormat       = "2006-01-02"
	FloatFormat                = "3.14"
	DefaultTimePipedriveFormat = "2006-01-02 15:04:05"
)

// POST CMS, COMMUNITY & EVENT
const (
	PostTypeCMS       = "cms"
	PostTypeArticle   = "article"
	PostTypeCommunity = "community"
)

const (
	PrimaryCalendarID = "c_bg5qdcajp11g29ttkkno395jp4@group.calendar.google.com"
)

const (
	SheetHotelOwner = "Hotel Owner"
	SheetWhiteLabel = "White Label"
	SheetGroup      = "Group"
	SheetBrand      = "Brand"
	SheetProperty   = "Property"
	SheetDepartment = "Department"
	SheetJob        = "Job"
)

// Activity Logging
const (
	DefaultMapSizeActivityLogging = 3

	ActivityLoggingFullName     = "fullName"
	ActivityLoggingTaskTitle    = "taskTitle"
	ActivityLoggingColumnTitle  = "columnTitle"
	ActivityLoggingColumnTitle2 = "columnTitle2"

	ActivityAddBoardTask     = "add-board-task"
	ActivityArchiveBoardTask = "archive-board-task"
	ActivityMoveBoardTask    = "move-board-task"
)

// Artist Satisfying Survey
const (
	SurveyTypeMultipleAnswer = "Multiple Answer"
	SurveyTypeEssay          = "Essay"
)

// Property Level
const (
	EnumPropertyLevelCollection    = "1"
	EnumPropertyLevelEconomy       = "2"
	EnumPropertyLevelMidScale      = "3"
	EnumPropertyLevelLifestyle     = "4"
	EnumPropertyLevelUpScale       = "5"
	EnumPropertyLevelFoodBeverages = "6"

	PropertyLevelCollection    = "Collections"
	PropertyLevelEconomy       = "Economy"
	PropertyLevelMidScale      = "Midscale"
	PropertyLevelLifestyle     = "Lifestyle"
	PropertyLevelUpScale       = "Upscale"
	PropertyLevelFoodBeverages = "Food & Beverages"
)

// Invitation Queue Status
const (
	InvitationQueueStatusQueued  = "queued"
	InvitationQueueStatusInvited = "invited"
)

// SQL Constraint
const (
	// Table Name
	TableEmployee = "employee"

	// Employee Unique Constraint
	UniqueNIKConstraint         = "unique_nik"
	UniqueIDCardConstraint      = "unique_id_card"
	UniqueNPWPConstraint        = "unique_npwp"
	UniqueEmailConstraint       = "unique_email"
	UniquePhoneNumberConstraint = "unique_phone_number"
)

const (
	QAStatusNew        = "new"
	QAStatusInProgress = "in_progress"
	QAStatusFinished   = "finished"

	QASubmissionStatusAvailable = "available"
	QASubmissionStatusFilled    = "filled"

	QAChecklistTypeScore  = "score"
	QAChecklistTypeNumber = "number"
	QAChecklistTypeText   = "text"

	QAValueC  = "C"
	QAValueNC = "NC"
	QAValueNA = "N/A"
)

// DepartmentGroup
const (
	EnumDepartmentUnitProperty = 1
	EnumDepartmentHeadquarters = 0

	DepartmentUnitProperty = "Unit Property"
	DepartmentHeadquarters = "Headquarters"
)

// Company Profile Group
const (
	CompanyProfileDocumentFlag      = "company profile"
	BrandProfileMorazenDocumentFlag = "brand profile morazen"
	BrandProfileLamoraDocumentFlag  = "brand profile lamora"
	BrandProfileOthersDocumentFlag  = "brand profile lainnya"
	ManualDocumentFlag              = "manual"
)

const (
	DC_BRI       ChannelCodeEnum = "DC_BRI"
	BCA_ONEKLIK  ChannelCodeEnum = "BCA_ONEKLIK"
	DC_MANDIRI   ChannelCodeEnum = "DC_MANDIRI"
	ID_OVO       ChannelCodeEnum = "ID_OVO"
	ID_DANA      ChannelCodeEnum = "ID_DANA"
	ID_LINKAJA   ChannelCodeEnum = "ID_LINKAJA"
	ID_SHOPEEPAY ChannelCodeEnum = "ID_SHOPEEPAY"
	ID_ASTRAPAY  ChannelCodeEnum = "ID_ASTRAPAY"
	ID_JENIUSPAY ChannelCodeEnum = "ID_JENIUSPAY"
	ID_SAKUKU    ChannelCodeEnum = "ID_SAKUKU"
)

var (
	DefaultExpiredTime       = 30
	DefaultExpiredQRTime     = 60
	SuccessRedirectURL       = "/payment-success"
	DefaultDynamicLink       = "DYNAMIC"
	DefaultOneTimePayment    = "ONE_TIME_PAYMENT"
	DefaultRedeemPoints      = "REDEEM_NONE"
	DefaultPrefixChannelCode = "ID_"
	DefaultCurrency          = "IDR"

	StatusSuccess     = "Confirm"
	StatusCodeSuccess = "00"

	// PaymentMethodCode
	PaymentMethodCodeVirtualAccount     = "VIRTUAL_ACCOUNT"
	PaymentMethodCodeQRIS               = "QRIS"
	PaymentMethodCodeEWallet            = "EWALLET"
	PaymentMethodCodeCreditCard         = "CREDIT_CARD"
	PaymentMethodCodeDirectDebit        = "DIRECT_DEBIT"
	PaymentMethodCodeVirtualAccountFlip = "virtual_account"
	PaymentMethodCodeEWalletFlip        = "wallet_account"

	// PaymentChannelCode
	PaymentChannelCodeDana      = "ID_DANA"
	PaymentChannelCodeOVO       = "ID_OVO"
	PaymentChannelCodeLinkAja   = "ID_LINKAJA"
	PaymentChannelCodeShopeePay = "ID_SHOPEEPAY"
	PaymentChannelCodeAstraPay  = "ID_ASTRAPAY"
	SourceTransaction           = "Transaction"
	PaymentMethodIDCreditCard   = "1141296b-eed2-4649-898d-d398750481ef"
	PaymentMethodIDDirectDebit  = "c0c73dfc-6d51-40cd-8676-a0c2580a8c40"

	PaymentGatewayXendit = "Xendit"
	PaymentGatewayFlip   = "Flip"
)

// Status Xendit
const (
	XenditStatusSucceeded = "SUCCEEDED"
	XenditStatusPending   = "PENDING"
	XenditStatusFailed    = "FAILED"
	XenditStatusVoided    = "VOIDED"
	XenditStatusRefunded  = "REFUNDED"
	XenditStatusCancelled = "Canceled"

	XenditStatusActive   = "ACTIVE"
	XenditStatusInactive = "INACTIVE"

	FlipStatusSucceeded = "SUCCESSFUL"
	FlipStatusFailed    = "FAILED"
	FlipStatusCanceled  = "CANCELLED"

	TransactionSubscription    = "Subscription"
	TransactionTopup           = "Topup"
	TransactionProduct         = "Product"
	TransactionTypeTransaction = "Transaction"
)

// Actor Update
const (
	DefaultActor       = `{"guid":"","name":"by system"}`
	DefaultActorSystem = "System"
)

const Charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const MembershipTypeDefault = "9d2a18c5-9048-434f-9b31-6a204462e72f"

// This is a constants for Customers service XENDIT purpose
const (
	ReusabilityMultipleUse = "MULTIPLE_USE"
	ReusabilityOneTimeUse  = "ONE_TIME_USE"
	DefaultTrue            = true
	DefaultFalse           = false
	CustomerTypeIndividual = "INDIVIDUAL"
	CustomerTypeBusiness   = "BUSINESS"
	NationalityIndonesia   = "ID"
	GenderMale             = "MALE"
	GenderFemale           = "FEMALE"
	GenderOther            = "OTHER"
	IdentifyAccountBA      = "BANK_ACCOUNT"
	IdentifyAccountEW      = "EWALLET"
	IdentifyAccountCC      = "CREDIT_CARD"
	IdentifyAccountQR      = "QR_CODE"
	CurrencyCodeIDR        = "IDR"
	CustomerDOBConstant    = "1990-01-01"
)

// This for redirect URL (credit cards)
const (
	SuccessCreditCardsURL = "https://www.google.com"
	FailureCreditCardsURL = "https://www.google.com"
)

// This a constants for xendit http request
const (
	RetryNumber   = 0
	TimeoutNumber = 20000
)

// This a constants for path URL
const (
	CustomerPathURL    = "/customers"
	TokensPathURL      = "/tokens"
	PaymentMethodV2URL = "/v2/payment_methods"
	PaymentRequestURL  = "/payment_requests"
	CreditCardsToken   = "/credit_card_tokens"
	CreditCardsCharges = "/credit_card_charges"
)

// This a constants for header request
const (
	HeaderAPIVersion   = "API-VERSION"
	CustomerAPIVersion = "2020-10-31"
)

// This constants for type of cards (CREDIT CARD or DEBIT CARD)
const (
	TypeCardCreditCard           = "CREDIT CARD"
	TypeCardDebitCard            = "DEBIT CARD"
	TypePaymentMethodCard        = "CARD"
	TypePaymentMethodDirectDebit = "DIRECT_DEBIT"
)

var (
	FirebaseNotificationTitle = "Mora Group"

	NotifCategoryIDTransaction  = "41359ac5-daa0-4152-848a-8a5502b3e2f3"
	NotifCategoryTransaction    = "Transaction"
	NotifCategoryIDSubscription = "ffe8d11d-4d2c-426d-9198-11a5087c511e"
	NotifCategorySubscription   = "Subscription"
	NotifCategoryIDTopup        = "2f043471-cf9b-4283-b0d9-ae22762cb512"
	NotifCategoryTopup          = "Topup"
	NotifCategoryIDBroadcast    = "f614ddad-4e7e-4ddb-9ff7-179be5c27c74"
	NotifCategoryBroadcast      = "Broadcast"
	NotifCategoryIDReminder     = "7a8c0a1d-32f6-4c7b-bc63-ffe4e23ba933"
	NotifCategoryReminder       = "Reminder"

	NotifActionAwaitingPaymentDetail     = "AWAITING_PAYMENT_DETAIL"
	NotifActionHistoryOrderDetail        = "HISTORY_ORDER_DETAIL"
	NotifActionHistorySubscriptionDetail = "HISTORY_SUBSCRIPTION_DETAIL"
	NotifActionHistoryRedeemDetail       = "HISTORY_REDEEM_DETAIL"
	NotifActionHistoryTopupDetail        = "HISTORY_TOPUP_DETAIL"
	NotifActionPostBroadcastDetail       = "POST BROADCAST"
)

// this is constants for doku
const (
	PaymentVirtualAccountBCA            = "VIRTUAL_ACCOUNT_BCA"
	PaymentVirtualAccountMandiri        = "VIRTUAL_ACCOUNT_BANK_MANDIRI"
	PaymentVirtualAccountSyariahMandiri = "VIRTUAL_ACCOUNT_BANK_SYARIAH_MANDIRI"
	PaymentVirtualAccountDoku           = "VIRTUAL_ACCOUNT_DOKU"
	PaymentVirtualAccountBRI            = "VIRTUAL_ACCOUNT_BRI"
	PaymentVirtualAccountBNI            = "VIRTUAL_ACCOUNT_BNI"
	PaymentVirtualAccountPermata        = "VIRTUAL_ACCOUNT_BANK_PERMATA"
	PaymentVirtualAccountCIMB           = "VIRTUAL_ACCOUNT_BANK_CIMB"
	PaymentVirtualAccountDanamon        = "VIRTUAL_ACCOUNT_BANK_DANAMON"
	PaymentOnlineToOfflineAlfa          = "ONLINE_TO_OFFLINE_ALFA"
	PaymentCreditCard                   = "CREDIT_CARD"
	PaymentDirectDebitBRI               = "DIRECT_DEBIT_BRI"
	PaymentEmoneyShopeePay              = "EMONEY_SHOPEEPAY"
	PaymentEmoneyOVO                    = "EMONEY_OVO"
	PaymentEmoneyDANA                   = "EMONEY_DANA"
	PaymentQRIS                         = "QRIS"
	PaymentPeerToPeerAkulaku            = "PEER_TO_PEER_AKULAKU"
	PaymentPeerToPeerKredivo            = "PEER_TO_PEER_KREDIVO"
	PaymentPeerToPeerIndodana           = "PEER_TO_PEER_INDODANA"

	PaymentDueDateDefault = 60

	PaymentGatewayDoku = "58178908-234b-42c0-8709-1107daf39178"
	PaymentMethodDoku  = "a56d48c0-65a8-4b5f-b9cc-625108aedbda"
)

var DokuDefaultTenor = []int{0, 3, 6, 12}

var PaymentMethodTypes = []string{
	PaymentVirtualAccountBCA,
	PaymentVirtualAccountMandiri,
	PaymentVirtualAccountSyariahMandiri,
	PaymentVirtualAccountDoku,
	PaymentVirtualAccountBRI,
	PaymentVirtualAccountBNI,
	PaymentVirtualAccountPermata,
	PaymentVirtualAccountCIMB,
	PaymentVirtualAccountDanamon,
	PaymentOnlineToOfflineAlfa,
	PaymentCreditCard,
	PaymentDirectDebitBRI,
	PaymentEmoneyShopeePay,
	PaymentEmoneyOVO,
	PaymentEmoneyDANA,
	PaymentQRIS,
	PaymentPeerToPeerAkulaku,
	PaymentPeerToPeerKredivo,
	PaymentPeerToPeerIndodana,
}

type PaymentMethodType string

const (
	PaymentMethodTypeCreditCard PaymentMethodType = "CREDIT_CARD"
)

type CategoryFirebaseNotification string

var (
	CategoryFirebaseNotificationSuccessCheckout     CategoryFirebaseNotification = "success_checkout"
	CategoryFirebaseNotificationSuccessPayment      CategoryFirebaseNotification = "success_payment"
	CategoryFirebaseNotificationStatusChanged       CategoryFirebaseNotification = "status_changed"
	CategoryFirebaseNotificationCanceledTransaction CategoryFirebaseNotification = "canceled"
	CategoryFirebaseNotificationExpiredTransaction  CategoryFirebaseNotification = "expired"
)

type NotificationMessageTransaction string

var (
	NotificationMessageTransactionSuccessCheckout          NotificationMessageTransaction = "Almost there! Your order is in, please finish your payment to complete the process."
	NotificationMessageTransactionPaymentSuccess           NotificationMessageTransaction = "Payment confirmed! 🚀 We're processing your order now. Stay tuned for updates."
	NotificationMessageTransactionTransactionStatusChanged NotificationMessageTransaction = "Heads up! Your transaction status has changed. Tap to view the latest update."
	NotificationMessageTransactionTransactionCanceled      NotificationMessageTransaction = "Your order has been canceled. Contact us if you need further assistance."
	NotificationMessageTransactionNewOrder                 NotificationMessageTransaction = "🎉 New order alert! You've just received a new order. Check it out now!"
	NotificationMessageTransactionReminder                 NotificationMessageTransaction = "Heads up! You’ve got a new order waiting to be processed. Don’t miss it!"
)

type XenditWebhookVerificationToken string

var (
	XenditWebhookVerificationTokenMorazenSurabaya XenditWebhookVerificationToken = "ZuzdrOdGYviN9VAqHWMo0IJU0XYUBQQXd0DQRYC88gkHYmby"
	// XenditWebhookVerificationTokenMorazenSurabaya   XenditWebhookVerificationToken = "TnSEfGSLg7w1Udp657Mm34CxNVnowHti8z14mB1XWgqvxmej"
	XenditWebhookVerificationTokenMorazenYogyakarta XenditWebhookVerificationToken = "kxTCLSVZJl0o1OfbVmpFNNTKhkUz1e1kU23RgSx3oXXzUSil" //prod
	// XenditWebhookVerificationTokenMorazenYogyakarta XenditWebhookVerificationToken = "GJpJNvm1jmSlhSZP28xd47wOdxJR43bT0u2DGF7irKbeXRLf" //dev
	XenditWebhookVerificationLamoraKotaLamaSurabaya XenditWebhookVerificationToken = "YqTeYulrzkXAPH8nuBLbY6GyKtSDB0y7UMuxg6Khb5z6htDa" //prod
	// XenditWebhookVerificationLamoraKotaLamaSurabaya XenditWebhookVerificationToken = "LTd4aaIRO8sCNgBUQKnETaJgqNNxEMIZNstZf445qIwGfhaX" //dev
	XenditWebhookVerificationLamoraSagan XenditWebhookVerificationToken = "JImVELUTKKVghJSK3HLtBstDCqIZNw44F8W6AmhJ41hMdQMO" //prod
)
