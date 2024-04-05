package csbalancing

import (
	"sort"
)

// Entity ...
type Entity struct {
	ID                     int
	Score                  int
	TotalCustomersAttended int
}

func getCustomerSuccessAvailable(customerSuccess []Entity, customerSuccessAway []int) []Entity {
	var customersAvaiable []Entity

	customersAway := make(map[int]bool)

	for _, id := range customerSuccessAway {
		customersAway[id] = true
	}

	for _, customer := range customerSuccess {
		if !customersAway[customer.ID] {
			customersAvaiable = append(customersAvaiable, customer)
		}
	}

	return customersAvaiable
}

func isAllCustomersScoresEqual(customers []Entity) bool {
	firstCustomerScore := customers[0].Score
	isAllScoresEqual := true

	for _, customer := range customers {
		if customer.Score != firstCustomerScore {
			isAllScoresEqual = false
			break
		}
	}

	return isAllScoresEqual
}

func sortCustomersByAscendingScore(customers []Entity) {
	compareFunc := func(x, y int) bool {
		return customers[x].Score < customers[y].Score
	}

	sort.Slice(customers, compareFunc)
}

func addTotalCustomersAttended(customers []Entity, customerSuccess []Entity) []Entity {

	for index, customerSuc := range customerSuccess {
		var customersAttended []Entity

		for _, customer := range customers {
			if customer.Score <= customerSuc.Score {
				customersAttended = append(customersAttended, customer)
			}
		}

		customers = customers[len(customersAttended):]

		customerSuccess[index].TotalCustomersAttended = len(customersAttended)

		if len(customers) == 0 {
			break
		}
	}

	return customerSuccess
}

func getCustomerSuccessWithMostCustomersAttended(customerSuccess []Entity) []Entity {

	var highestNumberOfCustomersAttended int

	for _, customerSuc := range customerSuccess {
		if customerSuc.TotalCustomersAttended > highestNumberOfCustomersAttended {
			highestNumberOfCustomersAttended = customerSuc.TotalCustomersAttended
		}
	}

	var result []Entity

	for _, customerSuc := range customerSuccess {
		if customerSuc.TotalCustomersAttended == highestNumberOfCustomersAttended {
			result = append(result, customerSuc)
		}
	}

	return result
}

func getCustomerSuccessIdWithMostCustomersAttended(customerSuccess []Entity) int {
	if len(customerSuccess) > 1 {
		return 0
	} else {
		return customerSuccess[0].ID
	}
}

// CustomerSuccessBalancing ...
func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {
	customerSuccessAvailable := getCustomerSuccessAvailable(customerSuccess, customerSuccessAway)

	if isAllCustomersScoresEqual(customers) {
		customers = []Entity{customers[0]}
	}

	sortCustomersByAscendingScore(customers)
	sortCustomersByAscendingScore(customerSuccessAvailable)

	customerSuccessWithTotalCustomersAttended := addTotalCustomersAttended(customers, customerSuccessAvailable)
	customerSuccessWithMostCustomersAttended := getCustomerSuccessWithMostCustomersAttended(customerSuccessWithTotalCustomersAttended)

	return getCustomerSuccessIdWithMostCustomersAttended(customerSuccessWithMostCustomersAttended)
}
