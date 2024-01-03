package print

type Icons struct {
	done     string
	looking  string
	fetching string
	linking  string
	building string
}

func UnicodeIcons() Icons {
	return Icons{
		done:     "✨",
		looking:  "🔍",
		fetching: "🚚",
		linking:  "🔗",
		building: "🔨",
	}
}
