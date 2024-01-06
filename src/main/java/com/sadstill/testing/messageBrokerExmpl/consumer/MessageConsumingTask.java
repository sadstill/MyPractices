package com.sadstill.testing.messageBrokerExmpl.consumer;

import com.sadstill.testing.messageBrokerExmpl.model.Message;
import com.sadstill.testing.messageBrokerExmpl.broker.MessageBroker;

import java.util.concurrent.TimeUnit;

public class MessageConsumingTask implements Runnable {
    private final MessageBroker messageBroker;

    private static final String TEMPLATE_MESSAGE_OF_MESSAGE_IS_CONSUMED = "Message '%s' is consumed. \n";

    public MessageConsumingTask(final MessageBroker messageBroker) {
        this.messageBroker = messageBroker;
    }

    @Override
    public void run() {
        while(!Thread.currentThread().isInterrupted()) {
            try {
                TimeUnit.SECONDS.sleep(1);
                final Message consumedMessage = this.messageBroker.consume();
                System.out.printf(TEMPLATE_MESSAGE_OF_MESSAGE_IS_CONSUMED, consumedMessage);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }
    }
}
