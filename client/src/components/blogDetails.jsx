import { useParams } from "react-router-dom";
import { Navbar } from "./Navbar";
import { useEffect, useState } from "react";
import axios from "axios";
import { FaSpinner } from 'react-icons/fa';

export const Details = () => {
    const { id } = useParams();
    const [Title, setTitle] = useState("");
    const [Post, setPost] = useState("");
    const [selectedFile, setSelectedFile] = useState(null);
    const [loading, setLoading] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const fetchData = async () => {
        try {
            setLoading(true);
            const response = await axios.get(`http://localhost:8000/blogdetail/${id}`);
            const content = response.data;
    
                setTitle(content.data.title);
                setPost(content.data.post);
                setSelectedFile(content.data.image);
                setLoading(false);
            } 
        catch (e) {
            console.error(e);
            setErrorMessage(e.message || "An error occurred"); 
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchData();
    }, [id]);

    if (loading) {
        return (
            <div className="spinner">
                <FaSpinner className="fa-spin" />
            </div>
        );
    }

    if (errorMessage) {
        return (
            <div className="error-message">
                Error Occurred: {errorMessage}
            </div>
        );
    }

    return (
        <>
            <Navbar />
            <div className="detail-container">
                <h2>Your post details</h2>
                <div className="post-detail-container">
                    <h3>{Title}</h3>
                    <p>{Post}</p>
                    {selectedFile && (
                        <img src={`http://localhost:8000${selectedFile}`} alt={Title} />
                    )}
                </div>
            </div>
        </>
    );
};
