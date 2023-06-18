package de.bit.workshop.moviedb.application;


import com.fasterxml.jackson.annotation.JsonInclude;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Builder
@Data
@NoArgsConstructor
@AllArgsConstructor
@JsonInclude(JsonInclude.Include.NON_NULL)
public class MovieDto {
    private String id;
    private String title;
    private String description;
    private List<String> actors;
    private int durationInMinutes;
    private String base64Cover;
}
