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

			connector, err := GetConnectorInstanceByName(msg.SurveyService.ConnectorName)
			if err != nil {
				contextLogger.Error(err.Error())
				continue
			}

			rows, err := connector.GetAnswers(msg.SurveyService.Name, msg.SurveyService.Address)
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
