import { IOneFilm, IFilms } from "../interfaces/App_interface";
import {BsFillHandThumbsUpFill, BsFillHandThumbsDownFill} from "react-icons/bs"
import axios, { AxiosResponse } from "axios";
import React, { FormEventHandler, useState } from "react";

export interface FilmProps {
  film: IOneFilm | undefined;
  getMovies(): any; 
}

const Film = ({ film, getMovies }: FilmProps) => {

  const [appComment, setAppComment] = useState('');
  
  const handleChange = (event: React.FormEvent<HTMLInputElement>) => {
    setAppComment(event.currentTarget.value);
  }

  const comments = (idFilm: number, comment: string) => {
    try {
      axios.put(`${process.env.REACT_APP_API_URL}/movie/comment`, {id: idFilm, comments: comment})
      .then((res: AxiosResponse<number>) => {
        getMovies();
      })
    } catch (error) {
      console.log(error);
    }
  }

  const removeComments = (idFilm: number) => {
    try {
      axios.put(`${process.env.REACT_APP_API_URL}/movie/de-comment`, {id: idFilm})
      .then((res: AxiosResponse<number>) => {
        getMovies();
      })
    } catch (error) {
      console.log(error);
    }
  }

  const likes = (idFilm: number) => {
    try {
      axios.put(`${process.env.REACT_APP_API_URL}/movie/like`, {id: idFilm})
      .then((res: AxiosResponse<number>) => {
        getMovies();
      })
    } catch (error) {
      console.log(error);
    }
  }

  const dislikes = (idFilm: number) => {
    try {
      axios.put(`${process.env.REACT_APP_API_URL}/movie/dislike`, {id: idFilm})
      .then((res: AxiosResponse<number>) => {
        getMovies();
      })
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <>
      {film?.map((value: IFilms, index: number, array: IOneFilm) => {

        return (
          <>
            <section
              className="film"
              key={index}>
              <div className="content">
                <h2>{value.title}</h2>
                <img src={`https://image.tmdb.org/t/p/w500${value.poster_path}`} alt="poster" />
                <p>{value.overview}</p>
              </div>
              <div className="likes-dislikes">
                <div className="likes">
                  <BsFillHandThumbsUpFill size={30} onClick={() => likes(value.id)}/>
                  <p>{value.likes}</p>
                </div>
                <div className="dislikes">
                  <BsFillHandThumbsDownFill size={30} onClick={() => dislikes(value.id)}/>
                  <p>{value.dislikes}</p>
                </div>
              </div>
              <div className="comments">
                <form className="comments-form" onSubmit={() => comments(value.id, appComment)}>
                  <input type="text" placeholder="Entrez un commentaire" onChange={handleChange}/>
                  <button type="submit">Ajouter</button>
                </form>
                {value.comments ? (
                  <div className="comments-display">
                    <p>{value.comments}</p>
                    <button onClick={() => removeComments(value.id)}>Supprimer</button>
                  </div>
                ) : (
                  <></>
                )}
              </div>
            </section>
          </>
        );
      })}
    </>
  );
};

export default Film;
