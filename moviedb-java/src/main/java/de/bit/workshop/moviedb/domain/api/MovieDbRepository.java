package de.bit.workshop.moviedb.domain.api;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

public interface MovieDbRepository {

    List<Movie> findAll();
    Optional<Movie> findById(UUID id);

    Movie createOrUpdate(Movie movie);

    void delete(UUID id);
}
