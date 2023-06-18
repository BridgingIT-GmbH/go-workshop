package de.bit.workshop.moviedb.domain.api;


import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;
import java.util.UUID;

@Builder
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Movie {
    private UUID id;
    private String title;
    private String description;
    private List<String> actors;
    private int durationInMinutes;
    private String base64Cover;
}
