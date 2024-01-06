package com.sadstill.testing.messageBrokerExmpl.producer;

import com.sadstill.testing.messageBrokerExmpl.broker.MessageBroker;
import com.sadstill.testing.messageBrokerExmpl.model.Message;

import java.util.concurrent.TimeUnit;

public final class MessageProducingTask implements Runnable {
    private static final String MESSAGE_PRODUCED_TO_BROKER = "Message '%s' is produced. \n";

    private final MessageBroker messageBroker;

    private final MessageFactory messageFactory;

    public MessageProducingTask(MessageBroker messageBroker) {
        this.messageBroker = messageBroker;
        this.messageFactory = new MessageFactory();
    }

    @Override
    public void run() {
        while(!Thread.currentThread().isInterrupted()) {
            try {
                final Message producedMessage = this.messageFactory.create();
                TimeUnit.SECONDS.sleep(3);
                this.messageBroker.produce(producedMessage);
                System.out.printf(MESSAGE_PRODUCED_TO_BROKER, producedMessage);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }
    }

    private static final class MessageFactory {
        private static final int INITIAL_NEXT_MESSAGE_INDEX = 1;

        private static final String TEMPLATE_CREATE_MESSAGE_DATA = "Message#%d";

        private int nextMessageIndex;

        public MessageFactory() {
            this.nextMessageIndex = INITIAL_NEXT_MESSAGE_INDEX;
        }

        public Message create() {
            return new Message(String.format(TEMPLATE_CREATE_MESSAGE_DATA, this.nextMessageIndex++));
        }
    }
}
