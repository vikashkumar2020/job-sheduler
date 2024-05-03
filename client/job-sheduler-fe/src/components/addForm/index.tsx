import { useState, ChangeEvent } from "react";
import { addNewJob } from "../../api/addNewJob";
import styles from "./index.module.css";

interface FormValues {
  name: string;
  duration: number;
}

function AddForm() {
  const [inputValue, setInputValue] = useState<FormValues>({ name: "", duration: 0 });
  const [error, setError] = useState<string>("");

  const onChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setInputValue(prevState => ({
      ...prevState,
      [name]: name === "duration" ? parseInt(value) : value
    }));
  };

  const validateInput = (values: FormValues): boolean => {
    values.name = values.name.trim()
    if (!values.name) {
      setError("Name is required");
      return false;
    }

    if (!values.duration || values.duration <= 0) {
      setError("Duration should be greater than 0");
      return false;
    }
    setError("");
    return true;
  };

  const addItem = () => {
    if (!validateInput(inputValue)) return;

    addNewJob(inputValue)
      .then(response => {
        console.log("Job added successfully:", response);
        setInputValue({ name: "", duration: 0 });
        setError("");
      })
      .catch(error => {
        console.error("Error adding job:", error);
        setError("Error adding job. Please try again later.");
      });
  };

  return (
    <div className={styles.container}>
    <label className={styles.label}>Name</label>
    <input type="text" name="name" value={inputValue.name} onChange={onChange} className={styles.input} />
    <label className={styles.label}>Duration</label>
    <input type="number" name="duration" value={inputValue.duration} onChange={onChange} className={styles.input} />
    {error && <p className={styles.error}>{error}</p>}
    <button onClick={addItem} className={styles.button}>Add</button>
  </div>
  );
}

export default AddForm;
