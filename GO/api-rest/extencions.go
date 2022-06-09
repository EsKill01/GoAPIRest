package main

func (this *MessageResponse) SetData(status string, message string) {
	this.Status = status
	this.Text = message
}
