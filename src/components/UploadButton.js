import React from "react";

function UploadButton({handleClick}) {
    return (
        < >
            <img src={require('../resources/upload.jpg')}
                onClick={handleClick} />
        </>
    );
}

export default UploadButton;