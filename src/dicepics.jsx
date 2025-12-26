const Dicepics = (props) => {
    if (props.value === 1) {
      return props.dicepic1
    } else if (props.value === 2) {
      return props.dicepic2
    } else if (props.value === 3) {
      return props.dicepic3
    } else if (props.value === 4) {
      return props.dicepic4
    } else if (props.value === 5) {
      return props.dicepic5
    } else if (props.value === 6) {
      return props.dicepic6
    } else {
      return props.dicepic0
    }
  }


export default Dicepics