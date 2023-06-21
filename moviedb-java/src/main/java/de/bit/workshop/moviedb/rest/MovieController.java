package de.bit.workshop.moviedb.rest;

import de.bit.workshop.moviedb.model.Movie;
import de.bit.workshop.moviedb.service.MovieService;
import lombok.AllArgsConstructor;
import org.springframework.data.rest.webmvc.ResourceNotFoundException;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping(value = "/movies", produces = "application/json")
@AllArgsConstructor
public class MovieController {

    private MovieService movieService;

    @GetMapping
    public List<Movie> getAll() {
        return movieService.loadAllMovies();
    }

    @GetMapping("/{id}")
    public Movie get(@PathVariable String id) {
        return movieService.loadMovieById(UUID.fromString(id)).orElseThrow(ResourceNotFoundException::new);
    }

    @GetMapping(value = "/{id}/cover")
    public String getCover(@PathVariable String id) {
        Movie movie = movieService.loadMovieById(UUID.fromString(id)).orElseThrow(ResourceNotFoundException::new);
        if (movie.getBase64Cover() == null) {
            throw new ResourceNotFoundException();
        }
        return movie.getBase64Cover();
    }

    @PostMapping
    public Movie createOrUpdate(@RequestBody Movie movie) {
        return movieService.createOrUpdateMovie(movie);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(value = HttpStatus.NO_CONTENT)
    public void delete(@PathVariable String id) {
        movieService.deleteMovie(UUID.fromString(id));
    }
}
