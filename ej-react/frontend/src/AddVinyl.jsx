import { useState} from "react";

const AddVinyl = () => {
    const [title, setTitle] = useState("");
    const [artist, setArtist] = useState("");
    const [year, setYear] = useState("");
    const [price, setPrice] = useState("");
    const [error, setError] = useState("");


    const handleCreate = async (e) => {
        e.preventDefault();
        setError("");

        try {
        const response = await fetch("http://localhost:8080/albums", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ title, artist, year: parseInt(year), price: parseFloat(price) }),
        });

        if (!response.ok) throw new Error("Error al tratar de crear un nuevo album");

        } catch {
            setError("Error al crear un nuevo album");
        }
    }

    return(
        <div>
            <form onSubmit={handleCreate}>
                <input
                    placeholder="Title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    required
                />
                <input
                    placeholder="Artist"
                    value={artist}
                    onChange={(e) => setArtist(e.target.value)}
                    required
                />
                <input
                    placeholder="Year"
                    value={year}
                    onChange={(e) => setYear(e.target.value)}
                />
                <input
                    placeholder="Price"
                    value={price}
                    onChange={(e) => setPrice(e.target.value)}
                />
                <button type="submit">Create</button>
                {error && <p style={{ color: "red" }}>{error}</p>}
            </form>
        </div>
    )

}

export default AddVinyl;