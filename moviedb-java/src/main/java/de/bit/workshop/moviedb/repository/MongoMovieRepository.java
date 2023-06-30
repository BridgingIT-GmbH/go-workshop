package de.bit.workshop.moviedb.repository;

import de.bit.workshop.moviedb.model.Movie;
import de.bit.workshop.moviedb.service.MovieRepository;
import lombok.AllArgsConstructor;
import org.springframework.context.annotation.Primary;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
@Primary
@AllArgsConstructor
public class MongoMovieRepository implements MovieRepository {

    private SpringDataMongoMovieRepository repo;
    @Override
    public List<Movie> findAll() {
        return repo.findAll();
    }

    @Override
    public Optional<Movie> findById(String id) {
        return repo.findById(id);
    }

    @Override
    public Movie createOrUpdate(Movie movie) {
        return repo.save(movie);
    }

    @Override
    public void delete(String id) {
        repo.deleteById(id);
    }
}
