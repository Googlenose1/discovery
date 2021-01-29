package api

// // GetSubscriptions ...
// func GetSubscriptions(c *gin.Context) {
// 	file, err := ioutil.ReadFile("data/filtered_subscriptions.json")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": fmt.Sprintf("Error reading file: %s", err.Error()),
// 		})
// 		return
// 	}

// 	var subscriptionList SubscriptionList
// 	err = json.Unmarshal(file, &subscriptionList)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": fmt.Sprintf("Error unmarshalling JSON: %s", err.Error()),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, subscriptionList)
// }
