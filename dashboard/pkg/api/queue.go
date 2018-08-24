package api

// QueueMessage tells which service to requests data from
type QueueMessage struct {
	SurveyService SurveyService
}

// StartQueueConsumer starts queue consumer
func (s *Server) StartQueueConsumer() {
	for {
		select {
		case msg := <-s.queue:
			contextLogger := s.logger.WithField("service", msg.SurveyService)
			contextLogger.Info("starting data fetch")

			connector := &SurvConnector{
				Name:    msg.SurveyService.Name,
				Address: msg.SurveyService.Address,
			}

			rows, err := connector.GetAnswers()
			if err != nil {
				contextLogger.WithError(err).Error("unable to fetch data")
				continue
			}

			for _, r := range rows {
				s.storage.Save(r)
			}
		}
	}
}
