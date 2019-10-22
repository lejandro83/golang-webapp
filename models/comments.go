package models

func GetComments() ([]string, error) {
	return client.LRange("comments", 0, 10).Result()
}

func PostComment(s string) error {
	return client.LPush("comments", s).Err()
}
