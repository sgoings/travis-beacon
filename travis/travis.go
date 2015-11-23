package travis

// WebhookPayload is the stuff
type WebhookPayload struct {
	Matrix []Job
}

// Job gives nice access to Job
type Job struct {
	Log string
}
