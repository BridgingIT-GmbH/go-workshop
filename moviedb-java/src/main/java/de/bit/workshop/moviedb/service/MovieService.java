package de.bit.workshop.moviedb.service;

import de.bit.workshop.moviedb.model.Movie;
import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
@AllArgsConstructor
public class MovieService {

    private MovieRepository rep;

    public List<Movie> loadAllMovies() {
        return rep.findAll();
    }

    public Optional<Movie> loadMovieById(UUID movieId) {
        return rep.findById(movieId);
    }

    public Movie createOrUpdateMovie(Movie movie) {
        if (movie.getId() == null) {
            movie.setId(UUID.randomUUID());
        }
        return rep.createOrUpdate(movie);
    }

    public void deleteMovie(UUID movieId) {
        rep.delete(movieId);
    }
}
