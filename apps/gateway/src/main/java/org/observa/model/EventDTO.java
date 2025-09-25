package org.observa.model;

import io.soabase.recordbuilder.core.RecordBuilder;

import java.time.Instant;
import java.util.List;
import java.util.Map;

@RecordBuilder
public record EventDTO(
        String eventId,
        String serviceName, String environment, Instant timestamp, String level,
        ErrorInfo error, Map<String, Object> context, Map<String, String> tags) {
    public static class ErrorInfo {
        public String type;
        public String message;
        public List<String> stackTrace;
    }
}

