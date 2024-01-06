package com.sadstill.testing.messageBrokerExmpl.broker;

import com.sadstill.testing.messageBrokerExmpl.model.Message;

import java.util.ArrayDeque;
import java.util.Queue;

public final class MessageBroker {
    private final Queue<Message> messagesToBeConsumed;

    private final int maxStoredMessages;

    public MessageBroker(final int maxStoredMessages) {
        this.messagesToBeConsumed = new ArrayDeque<>(maxStoredMessages);
        this.maxStoredMessages = maxStoredMessages;
    }

    public synchronized void produce(final Message message) {

        while (this.messagesToBeConsumed.size() >= this.maxStoredMessages) {
            try {
                wait();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }

        messagesToBeConsumed.add(message);
        notify();
    }

    public synchronized Message consume() {

        while (this.messagesToBeConsumed.isEmpty()) {
            try {
                wait();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }

        Message consumedMessage = this.messagesToBeConsumed.poll();
        notify();
        return consumedMessage;
    }
}
