
interface newJobProps{
  name:string;
  duration: number;
}
export async function addNewJob(newJob:newJobProps) {
  
  try {
    let response=await fetch(`http://${import.meta.env.VITE_BACKEND_URL}/job`,{
      method:'POST',
      headers:{
        "Content-Type": "application/json",
      },
      body:JSON.stringify(newJob)
    });
    let data= await response.json();
    return data;
  } catch (error) {
    return error
  }
}