package service

var (
	GET_ISSUES = map[string]string{
		"query": `{
			repository(name: "%s", owner: "%s") {
					issues (states: CLOSED, first: 100) {
							nodes {
									title
									assignees (first: 100) {
											nodes {
													name
											}
									}
							}
					}
			}
		}`}
)
