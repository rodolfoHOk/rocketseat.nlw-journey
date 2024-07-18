package br.com.rocketseat.hiokdev.planner_java.api.trip.dto;

import br.com.rocketseat.hiokdev.planner_java.config.validation.ValidDateTime;
import br.com.rocketseat.hiokdev.planner_java.domain.trip.Trip;
import io.swagger.v3.oas.annotations.media.Schema;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public record TripUpdateRequestPayload(
        @Schema(example = "Porto Alegre, RS") @NotBlank @Size(min = 4, max = 255) String destination,
        @Schema(example = "2024-07-25T10:30:00Z") @NotBlank @ValidDateTime String starts_at,
        @Schema(example = "2024-07-29T12:00:00Z") @NotBlank @ValidDateTime String ends_at
) {

    public static Trip toDomain(TripUpdateRequestPayload payload) {
        return Trip.builder()
                .destination(payload.destination)
                .startsAt(LocalDateTime.parse(payload.starts_at(), DateTimeFormatter.ISO_DATE_TIME))
                .endsAt(LocalDateTime.parse(payload.ends_at(), DateTimeFormatter.ISO_DATE_TIME))
                .build();
    }

}
