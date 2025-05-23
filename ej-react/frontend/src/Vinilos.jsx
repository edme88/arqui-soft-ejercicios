import { useEffect, useState } from "react";

const Vinilos = () => {
    const [vinilos, setVinilos] = useState([]);

    useEffect(() => {
        fetch("http://localhost:8080/albums")
        .then((res) => res.json())
        .then((data) => setVinilos(data))
        .catch((err) => console.error("Error fetching albums: ", err))
    }, []);

    return (
        <div className="vinilos-container">
            { vinilos.map((album, index) => (
                <div className="card-album" key={index}>
                    <h2>{album.title}</h2>
                    <h2>{album.artist}</h2>
                    <h2>{album.year}</h2>
                    <h2>{album.price}</h2>
                </div>
            ))}  
        </div>
    )

}

export default Vinilos;