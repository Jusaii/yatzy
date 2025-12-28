import { useState } from "react";
import { updateSubTotal, perItemScores } from "../totals";

function Numberbuttons(props) {
  const [buttonText, setButtonText] = useState(props.text);
  const [isClicked, setIsClicked] = useState(false);

  const handleClick = () => {
    const values = props.values

    const targetNumber = props.target;
    const count = values.filter(value => value === targetNumber).length;
    const sum = props.target * count;

    perItemScores.set(`${targetNumber}`, sum)
    console.log(`Sum of ${targetNumber}'s: ${perItemScores.get(`${targetNumber}`)}`)
    setButtonText(`${sum}`);
    setIsClicked(true);
    updateSubTotal(sum);
    props.reset()
  }

  return (
    <button
      className="select-container"
      onClick={handleClick}
      disabled={isClicked}
    >
      {buttonText}
    </button>
  );
}

export default Numberbuttons
