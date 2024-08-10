import axios from "axios"
import { useState} from "react"
import {FaSpinner } from 'react-icons/fa';
import { useNavigate, useParams } from "react-router-dom";
import { Navbar } from "./Navbar";

export const DeleteBlog = ()=>{
    const [message,setMessage] = useState("")
    const [loading,setLoading] = useState(false)
    const [showmessage,setshowmessage] = useState(false)
    const navigate = useNavigate()
    const {id} = useParams()

    const handleDelete = async()=>{
        try {
            setLoading(true)
          const response = await axios.delete(`http://localhost:8000/deleteblog/${id}`)
          setMessage(response.data.message);
          setLoading(false)
          setshowmessage(true)
        } catch (error) {
          console.error("Error occurred while deleting blog:", error);
          setMessage(error.response?.data?.message || "Something went wrong");
          setLoading(false)
          setshowmessage(false)
        }
    }
    if (loading) {
        return (
          <div className="spinner">
            <FaSpinner className="fa-spin" />
          </div>
        );
      }
      return (
        <>
       <Navbar/>
        <div className="topContainer">
        <div className="delete-prompt-container">
            <h3>Are you sure you want to delete this post</h3>
            <button onClick={handleDelete}>Yes</button>
            <button onClick={()=>navigate("/")}>Cancel</button>
            {showmessage &&  <p>{message}</p>}
        </div>
        </div>
        </>
      )

}