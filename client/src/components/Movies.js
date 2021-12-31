import React, { Component, Fragment } from 'react'
import { Link } from 'react-router-dom'

export default class Movies extends Component {

    state = {
        movies: [],
        isLoaded: false,
        error: null
    }

    componentDidMount() {
        fetch("http://localhost:4000/v1/movies")
            .then((response) => {
                if (response.status !== "200") {
                    let err = Error
                    err.Message = "Invalid response code: " + response.status
                    this.setState({ error: err })
                }
                return response.json()
            })
            .then((json) => {
                this.setState({
                    movies: json.movies,
                    isLoaded: true,
                },
                    (error) => {
                        this.setState({
                            isLoaded: true,
                            error
                        })
                    }
                )
            })
    }

    render() {
        const { movies, isLoaded, error } = this.state
        if (error) {
            return <div>Error: {error.Message}</div>
        } else if (!isLoaded) {
            return <p>Loading...</p>
        } else {
            return (
                <Fragment>
                    <h2>Choose a movie</h2>
                    <ul>
                        {movies.map(movie => (
                            <li key={movie.id}>
                                <Link to={`/movies/${movie.id}`}>{movie.title}</Link>
                            </li>
                        ))}
                    </ul>
                </Fragment>
            )
        }
    }
}
