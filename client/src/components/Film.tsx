import React from "react";
import { IOneFilm, IFilms } from "../interfaces/App_interface";

export interface FilmProps {
  film: IOneFilm | undefined;
}

const Film = ({ film }: FilmProps) => {
  return (
    <>
      {film?.map((value: IFilms, index: number, array: IOneFilm) => {
        console.log(value);
        console.log(index);
        console.log(array);
        return (
          <main key={index}>
            <h2>{value.Title}</h2>
            <img src={value.Image} alt="Affiche du film" />
            <p>{value.Description}</p>
          </main>
        );
      })}
    </>
  );
};

export default Film;
