package de.bit.workshop.moviedb.model;


import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Builder
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Movie {
    private String id;
    private String title;
    private String description;
    private List<String> actors;
    private int durationInMinutes;
    @JsonProperty(access = JsonProperty.Access.WRITE_ONLY)
    private String base64Cover;
}
