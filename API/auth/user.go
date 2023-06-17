package auth

type (
	User struct {
		View                            string         `json:"view"`
		Avatar                          any            `json:"avatar"`
		AvatarThumbs                    any            `json:"avatarThumbs"`
		Header                          any            `json:"header"`
		HeaderSize                      any            `json:"headerSize"`
		HeaderThumbs                    any            `json:"headerThumbs"`
		Id                              int            `json:"id"`
		Name                            string         `json:"name"`
		Username                        string         `json:"username"`
		CanLookStory                    bool           `json:"canLookStory"`
		CanCommentStory                 bool           `json:"canCommentStory"`
		HasNotViewedStory               bool           `json:"hasNotViewedStory"`
		IsVerified                      bool           `json:"isVerified"`
		CanPayInternal                  bool           `json:"canPayInternal"`
		HasScheduledStream              bool           `json:"hasScheduledStream"`
		HasStream                       bool           `json:"hasStream"`
		HasStories                      bool           `json:"hasStories"`
		TipsEnabled                     bool           `json:"tipsEnabled"`
		TipsTextEnabled                 bool           `json:"tipsTextEnabled"`
		TipsMin                         int            `json:"tipsMin"`
		TipsMax                         int            `json:"tipsMax"`
		CanEarn                         bool           `json:"canEarn"`
		CanAddSubscriber                bool           `json:"canAddSubscriber"`
		SubscribePrice                  int            `json:"subscribePrice"`
		HasStripe                       bool           `json:"hasStripe"`
		IsStripeExist                   bool           `json:"isStripeExist"`
		SubscriptionBundles             []any          `json:"subscriptionBundles"`
		CanSendChatToAll                bool           `json:"canSendChatToAll"`
		CreditsMin                      int            `json:"creditsMin"`
		CreditsMax                      int            `json:"creditsMax"`
		IsPaywallRestriction            bool           `json:"isPaywallRestriction"`
		Unprofitable                    bool           `json:"unprofitable"`
		ListsSort                       string         `json:"listsSort"`
		ListsSortOrder                  string         `json:"listsSortOrder"`
		CanCreateLists                  bool           `json:"canCreateLists"`
		JoinDate                        string         `json:"joinDate"`
		IsReferrerAllowed               bool           `json:"isReferrerAllowed"`
		About                           string         `json:"about"`
		RawAbout                        string         `json:"rawAbout"`
		Website                         string         `json:"website"`
		Wishlist                        string         `json:"wishlist"`
		Location                        string         `json:"location"`
		PostsCount                      int            `json:"postsCount"`
		ArchivedPostsCount              int            `json:"archivedPostsCount"`
		PhotosCount                     int            `json:"photosCount"`
		VideosCount                     int            `json:"videosCount"`
		AudiosCount                     int            `json:"audiosCount"`
		MediasCount                     int            `json:"mediasCount"`
		Promotions                      []any          `json:"promotions"`
		LastSeen                        any            `json:"lastSeen"`
		FavoritesCount                  int            `json:"favoritesCount"`
		ShowPostsInFeed                 bool           `json:"showPostsInFeed"`
		CanReceiveChatMessage           bool           `json:"canReceiveChatMessage"`
		IsPerformer                     bool           `json:"isPerformer"`
		IsRealPerformer                 bool           `json:"isRealPerformer"`
		IsSpotifyConnected              bool           `json:"isSpotifyConnected"`
		SubscribersCount                int            `json:"subscribersCount"`
		HasPinnedPosts                  bool           `json:"hasPinnedPosts"`
		CanChat                         bool           `json:"canChat"`
		CallPrice                       int            `json:"callPrice"`
		IsPrivateRestriction            bool           `json:"isPrivateRestriction"`
		ShowSubscribersCount            bool           `json:"showSubscribersCount"`
		ShowMediaCount                  bool           `json:"showMediaCount"`
		SubscribedByData                any            `json:"subscribedByData"`
		SubscribedOnData                any            `json:"subscribedOnData"`
		CanPromotion                    bool           `json:"canPromotion"`
		CanCreatePromotion              bool           `json:"canCreatePromotion"`
		CanCreateTrial                  bool           `json:"canCreateTrial"`
		IsAdultContent                  bool           `json:"isAdultContent"`
		IsBlocked                       bool           `json:"isBlocked"`
		CanTrialSend                    bool           `json:"canTrialSend"`
		CanAddPhone                     bool           `json:"canAddPhone"`
		PhoneLast4                      any            `json:"phoneLast4"`
		PhoneMask                       any            `json:"phoneMask"`
		HasNewTicketReplies             map[string]any `json:"hasNewTicketReplies"`
		HasInternalPayments             bool           `json:"hasInternalPayments"`
		IsCreditsEnabled                bool           `json:"isCreditsEnabled"`
		CreditBalance                   float64        `json:"creditBalance"`
		IsMakePayment                   bool           `json:"isMakePayment"`
		IsOtpEnabled                    bool           `json:"isOtpEnabled"`
		Email                           string         `json:"email"`
		IsEmailChecked                  bool           `json:"isEmailChecked"`
		IsLegalApprovedAllowed          bool           `json:"isLegalApprovedAllowed"`
		IsTwitterConnected              bool           `json:"isTwitterConnected"`
		TwitterUsername                 any            `json:"twitterUsername"`
		IsAllowTweets                   bool           `json:"isAllowTweets"`
		IsPaymentCardConnected          bool           `json:"isPaymentCardConnected"`
		ReferalUrl                      string         `json:"referalUrl"`
		IsVisibleOnline                 bool           `json:"isVisibleOnline"`
		SubscribesCount                 int            `json:"subscribesCount"`
		CanPinPost                      bool           `json:"canPinPost"`
		HasNewAlerts                    bool           `json:"hasNewAlerts"`
		HasNewHints                     bool           `json:"hasNewHints"`
		HasNewChangedPriceSubscriptions bool           `json:"hasNewChangedPriceSubscriptions"`
		NotificationsCount              int            `json:"notificationsCount"`
		ChatMessagesCount               int            `json:"chatMessagesCount"`
		IsWantComments                  bool           `json:"isWantComments"`
		WatermarkText                   string         `json:"watermarkText"`
		CustomWatermarkText             any            `json:"customWatermarkText"`
		HasWatermarkPhoto               bool           `json:"hasWatermarkPhoto"`
		HasWatermarkVideo               bool           `json:"hasWatermarkVideo"`
		HanDelete                       bool           `json:"canDelete"`
		IsTelegramConnected             bool           `json:"isTelegramConnected"`
		AdvBlock                        []any          `json:"advBlock"`
		HasPurchasedPosts               bool           `json:"hasPurchasedPosts"`
		IsEmailRequired                 bool           `json:"isEmailRequired"`
		IsPayoutLegalApproved           bool           `json:"isPayoutLegalApproved"`
		PayoutLegalApproveState         string         `json:"payoutLegalApproveState"`
		PayoutLegalApproveRejectReason  any            `json:"payoutLegalApproveRejectReason"`
		EnabledImageEditorForChat       bool           `json:"enabledImageEditorForChat"`
		ShouldReceiveLessNotifications  bool           `json:"shouldReceiveLessNotifications"`
		CanCalling                      bool           `json:"canCalling"`
		PaidFeed                        bool           `json:"paidFeed"`
		CanSendSms                      bool           `json:"canSendSms"`
		CanAddFriends                   bool           `json:"canAddFriends"`
		IsRealCardConnected             bool           `json:"isRealCardConnected"`
		CountPriorityChat               int            `json:"countPriorityChat"`
		HasScenario                     bool           `json:"hasScenario"`
		IsWalletAutorecharge            bool           `json:"isWalletAutorecharge"`
		WalletAutorechargeAmount        int            `json:"walletAutorechargeAmount"`
		WalletAutorechargeMin           int            `json:"walletAutorechargeMin"`
		WalletFirstRebills              bool           `json:"walletFirstRebills"`
		CloseFriends                    int            `json:"closeFriends"`
		CanAlternativeWalletTopUp       bool           `json:"canAlternativeWalletTopUp"`
		NeedIVApprove                   bool           `json:"needIVApprove"`
		IvStatus                        any            `json:"ivStatus"`
		IvFailReason                    any            `json:"ivFailReason"`
		CanCheckDocsOnAddCard           bool           `json:"canCheckDocsOnAddCard"`
		FaceIdAvailable                 bool           `json:"faceIdAvailable"`
		IvCountry                       any            `json:"ivCountry"`
		IvForcedVerified                bool           `json:"ivForcedVerified"`
		IvFlow                          string         `json:"ivFlow"`
		IsVerifiedReason                bool           `json:"isVerifiedReason"`
		CanReceiveManualPayout          bool           `json:"canReceiveManualPayout"`
		CanReceiveStripePayout          bool           `json:"canReceiveStripePayout"`
		ManualPayoutPendingDays         int            `json:"manualPayoutPendingDays"`
		IsNeedConfirmPayout             bool           `json:"isNeedConfirmPayout"`
		CanStreaming                    bool           `json:"canStreaming"`
		IsScheduledStreamsAllowed       bool           `json:"isScheduledStreamsAllowed"`
		CanMakeExpirePosts              bool           `json:"canMakeExpirePosts"`
		TrialMaxDays                    int            `json:"trialMaxDays"`
		TrialMaxExpiresDays             int            `json:"trialMaxExpiresDays"`
		MessageMinPrice                 int            `json:"messageMinPrice"`
		MessageMaxPrice                 int            `json:"messageMaxPrice"`
		PostMinPrice                    int            `json:"postMinPrice"`
		PostMaxPrice                    int            `json:"postMaxPrice"`
		StreamMinPrice                  int            `json:"streamMinPrice"`
		StreamMaxPrice                  int            `json:"streamMaxPrice"`
		CanCreatePaidStream             bool           `json:"canCreatePaidStream"`
		CallMinPrice                    int            `json:"callMinPrice"`
		CallMaxPrice                    int            `json:"callMaxPrice"`
		SubscribeMinPrice               float64        `json:"subscribeMinPrice"`
		SubscribeMaxPrice               int            `json:"subscribeMaxPrice"`
		BundleMaxPrice                  int            `json:"bundleMaxPrice"`
		UnclaimedOffersCount            int            `json:"unclaimedOffersCount"`
		ClaimedOffersCount              int            `json:"claimedOffersCount"`
		WithdrawalPeriod                string         `json:"withdrawalPeriod"`
		CanAddStory                     bool           `json:"canAddStory"`
		CanAddSubscriberByBundle        bool           `json:"canAddSubscriberByBundle"`
		IsSuggestionsOptOut             bool           `json:"isSuggestionsOptOut"`
		CanCreateFundRaising            bool           `json:"canCreateFundRaising"`
		MinFundRaisingTarget            int            `json:"minFundRaisingTarget"`
		MaxFundRaisingTarget            int            `json:"maxFundRaisingTarget"`
		DisputesRatio                   int            `json:"disputesRatio"`
		VaultListsSort                  string         `json:"vaultListsSort"`
		VaultListsSortOrder             string         `json:"vaultListsSortOrder"`
		CanCreateVaultLists             bool           `json:"canCreateVaultLists"`
		CanMakeProfileLinks             bool           `json:"canMakeProfileLinks"`
		ReplyOnSubscribe                bool           `json:"replyOnSubscribe"`
		PayoutType                      string         `json:"payoutType"`
		MinPayoutSumm                   int            `json:"minPayoutSumm"`
		CanHasW9Form                    bool           `json:"canHasW9Form"`
		IsVatRequired                   bool           `json:"isVatRequired"`
		IsCountryVatRefundable          bool           `json:"isCountryVatRefundable"`
		IsCountryVatNumberCollect       bool           `json:"isCountryVatNumberCollect"`
		VatNumberName                   string         `json:"vatNumberName"`
		IsCountryWithVat                bool           `json:"isCountryWithVat"`
		ConnectedOfAccounts             []any          `json:"connectedOfAccounts"`
		HasPassword                     bool           `json:"hasPassword"`
		CanConnectOfAccount             bool           `json:"canConnectOfAccount"`
		PinnedPostsCount                int            `json:"pinnedPostsCount"`
		MaxPinnedPostsCount             int            `json:"maxPinnedPostsCount"`
	}
	Visibility struct {
		IsVisibleOnline bool `json:"isVisibleOnline"`
	}
)