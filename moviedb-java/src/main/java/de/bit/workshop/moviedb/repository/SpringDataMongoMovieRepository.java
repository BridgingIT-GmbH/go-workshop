package de.bit.workshop.moviedb.repository;

import de.bit.workshop.moviedb.model.Movie;
import org.bson.types.ObjectId;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface SpringDataMongoMovieRepository extends MongoRepository<Movie, String> {
}
