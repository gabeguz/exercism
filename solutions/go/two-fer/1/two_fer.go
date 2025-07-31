// Package twofer allows you to share something with someone.
package twofer

// ShareWith accetps a string (name) and prints a message with that name, or a
// generic message if no name is provided.
func ShareWith(name string) string {
	if len(name) > 0 {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}