package com.sadstill.testing.messageBrokerExmpl;

import com.sadstill.testing.messageBrokerExmpl.consumer.MessageConsumingTask;
import com.sadstill.testing.messageBrokerExmpl.broker.MessageBroker;
import com.sadstill.testing.messageBrokerExmpl.producer.MessageProducingTask;

public class Runner {
    public static void main(String[] args) {
        final int brokerMaxStoredMessage = 5;
        final MessageBroker messageBroker = new MessageBroker(brokerMaxStoredMessage);

        final Thread producingThread = new Thread(new MessageProducingTask(messageBroker));
        final Thread consumingThread = new Thread(new MessageConsumingTask(messageBroker));

        producingThread.start();
        consumingThread.start();
    }
}
