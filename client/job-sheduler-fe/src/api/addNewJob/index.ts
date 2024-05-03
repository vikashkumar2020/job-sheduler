
interface newJobProps{
  name:string;
  duration: number;
}
export async function addNewJob(newJob:newJobProps) {
  
  try {
    let response=await fetch(`http://localhost:8080/api/v1/job`,{
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