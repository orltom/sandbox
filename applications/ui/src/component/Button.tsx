import './Button.css';

interface Props {
    text: string;
    float: string
}

function Button(props: Props) {
    return (
        <div className={`navigation-button ${props.float == "left" ? "navigation-button-left" : "navigation-button-right"}`} >{props.text}</div>
    );
}

export default Button;
