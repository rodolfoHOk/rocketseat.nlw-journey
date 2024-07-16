package br.com.rocketseat.hiokdev.planner_java.api.trip.dto;

import br.com.rocketseat.hiokdev.planner_java.domain.trip.Trip;
import lombok.Builder;

import java.time.LocalDateTime;
import java.util.UUID;

public record TripData(
        UUID id,
        String destination,
        LocalDateTime startsAt,
        LocalDateTime endsAt,
        Boolean isConfirmed
) {

    @Builder
    public TripData {
    }

    public static TripData toResponse(Trip trip) {
        return TripData.builder()
                .id(trip.getId())
                .destination(trip.getDestination())
                .startsAt(trip.getStartsAt())
                .endsAt(trip.getEndsAt())
                .isConfirmed(trip.getIsConfirmed())
                .build();
    }

}