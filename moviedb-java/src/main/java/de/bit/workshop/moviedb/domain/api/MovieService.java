package de.bit.workshop.moviedb.domain.api;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

public interface MovieService {

    List<Movie> loadAllMovies();
    Optional<Movie> loadMovieById(UUID movieId);
    Movie createOrUpdateMovie(Movie movie);
    void deleteMovie(UUID movieId);
}
