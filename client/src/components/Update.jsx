import axios from "axios"
import {useState } from "react"
import { Navbar } from "./Navbar"
import {FaSpinner } from 'react-icons/fa';
import { useParams } from "react-router-dom";

export const UpdateBlog = ()=>{
    const [Title,setTitle] = useState("")
    const [Post,setPost] = useState("")
    const [selectedFile,setSelectedFile] = useState(null)
    const [message,setMessage] = useState("")
    const [loading,setLoading] = useState(false)
    const {id} = useParams()

    const handleSubmit = async (e) => {
      e.preventDefault();
    
      const formData = new FormData();
      formData.append("title", Title);
      formData.append("post", Post);
      if (selectedFile) {
        formData.append("image", selectedFile);
      }
      try {
          setLoading(true)
        const response = await axios.put(`http://localhost:8000/updateblog/${id}`, formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });
        setMessage(response.data.message);
        setLoading(false)
      } catch (error) {
        console.error("Error occurred while creating blog:", error);
        setMessage(error.response?.data?.message || "Something went wrong");
        setLoading(false)
      }
    };

      if (loading) {
        return (
          <div className="spinner">
            <FaSpinner className="fa-spin" />
          </div>
        );
      }

    return(   
        <>
        <Navbar/>
        <div className="form-Container">
        <h2>Update your post</h2>
            <form className="form-cls" onSubmit={handleSubmit}>
                <input required type="text" placeholder="Enter Title" onChange={(e)=>setTitle(e.target.value)}></input>
                <input required type="text" placeholder="Enter Post Content" onChange={(e)=>setPost(e.target.value)}></input>
                <input type="file" onChange={(e)=>setSelectedFile(e.target.files[0])}></input>
                <input type="Submit"></input>
            </form>
            {message && <h3>{message}</h3>}
        </div>
        </>
    
    )
}