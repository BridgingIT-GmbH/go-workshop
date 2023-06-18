package de.bit.workshop.moviedb.application;

import de.bit.workshop.moviedb.domain.api.Movie;
import de.bit.workshop.moviedb.domain.api.MovieDbService;
import jakarta.annotation.PostConstruct;
import lombok.AllArgsConstructor;
import org.modelmapper.ModelMapper;
import org.modelmapper.config.Configuration;
import org.springframework.data.rest.webmvc.ResourceNotFoundException;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;
import java.util.stream.Collectors;

@RestController
@RequestMapping(value = "/movies", produces = "application/json")
@AllArgsConstructor
public class MovieDbController {

    private MovieDbService movieDbService;
    private final ModelMapper modelMapper = new ModelMapper();

    @PostConstruct
    public void init() {
        modelMapper.getConfiguration()
                .setFieldMatchingEnabled(true)
                .setFieldAccessLevel(Configuration.AccessLevel.PRIVATE);
    }

    @GetMapping
    public List<MovieDto> getAll() {
        List<Movie> movies = movieDbService.loadAllMovies();
        return mapToDto(movies);
    }

    @GetMapping("/{id}")
    public MovieDto get(@PathVariable String id) {
        Movie movie = movieDbService.loadMovieById(UUID.fromString(id)).orElseThrow(ResourceNotFoundException::new);
        return this.mapToDto(movie);
    }

    @GetMapping(value = "/{id}/cover")
    public String getCover(@PathVariable String id) {
        Movie movie = movieDbService.loadMovieById(UUID.fromString(id)).orElseThrow(ResourceNotFoundException::new);
        if (movie.getBase64Cover() == null) {
            throw new ResourceNotFoundException();
        }
        return movie.getBase64Cover();
    }

    @PostMapping
    public MovieDto createOrUpdate(@RequestBody MovieDto movieDto) {
        Movie movie = movieDbService.createOrUpdateMovie(modelMapper.map(movieDto, Movie.class));
        return this.mapToDto(movie);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(value = HttpStatus.NO_CONTENT)
    public void delete(@PathVariable String id) {
        movieDbService.deleteMovie(UUID.fromString(id));
    }

    private MovieDto mapToDto(Movie movie) {
        MovieDto movieDto = modelMapper.map(movie, MovieDto.class);
        movieDto.setBase64Cover(null);
        return movieDto;
    }

    private List<MovieDto> mapToDto(List<Movie> movies) {
        return movies.stream().map(this::mapToDto).collect(Collectors.toList());
    }
}
