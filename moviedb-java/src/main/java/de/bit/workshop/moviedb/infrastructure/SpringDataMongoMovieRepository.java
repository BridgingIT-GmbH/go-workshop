package de.bit.workshop.moviedb.infrastructure;

import de.bit.workshop.moviedb.domain.api.Movie;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface SpringDataMongoMovieRepository extends MongoRepository<Movie, UUID> {
}
