package de.bit.workshop.moviedb.infrastructure;

import de.bit.workshop.moviedb.domain.api.Movie;
import de.bit.workshop.moviedb.domain.api.MovieDbRepository;
import lombok.AllArgsConstructor;
import org.springframework.context.annotation.Primary;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
@Primary
@AllArgsConstructor
public class MongoMovieDbRepository implements MovieDbRepository {

    private SpringDataMongoMovieDbRepository repo;
    @Override
    public List<Movie> findAll() {
        return repo.findAll();
    }

    @Override
    public Optional<Movie> findById(UUID id) {
        return repo.findById(id);
    }

    @Override
    public Movie createOrUpdate(Movie movie) {
        return repo.save(movie);
    }

    @Override
    public void delete(UUID id) {
        repo.deleteById(id);
    }
}
