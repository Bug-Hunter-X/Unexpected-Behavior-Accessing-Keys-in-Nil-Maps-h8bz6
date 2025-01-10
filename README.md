# Unexpected Behavior Accessing Keys in Nil Maps

This example demonstrates the unexpected behavior when accessing keys in nil maps in Go.  In many other languages, this would result in a runtime error.  However, in Go, it silently returns the zero value for the map's value type.

The solution shows how to check for nil maps before accessing keys to prevent unexpected results and potential bugs.

This is a common pitfall for developers transitioning from other languages to Go.  Always check if the map is initialized before accessing keys to avoid this issue.
