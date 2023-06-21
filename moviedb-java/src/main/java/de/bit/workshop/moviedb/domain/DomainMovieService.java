package de.bit.workshop.moviedb.domain;

import de.bit.workshop.moviedb.domain.api.Movie;
import de.bit.workshop.moviedb.domain.api.MovieRepository;
import de.bit.workshop.moviedb.domain.api.MovieService;
import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
@AllArgsConstructor
public class DomainMovieService implements MovieService {

    private MovieRepository rep;

    @Override
    public List<Movie> loadAllMovies() {
        return rep.findAll();
    }

    @Override
    public Optional<Movie> loadMovieById(UUID movieId) {
        return rep.findById(movieId);
    }

    @Override
    public Movie createOrUpdateMovie(Movie movie) {
        if (movie.getId() == null) {
            movie.setId(UUID.randomUUID());
        }
        return rep.createOrUpdate(movie);
    }

    @Override
    public void deleteMovie(UUID movieId) {
        rep.delete(movieId);
    }
}
