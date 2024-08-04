//For setting key-value pairs and expiration time.

import React, { useState } from 'react';
import axios from 'axios';
import "./CacheInput.css"

const CacheInput = () => {
    const [key, setKey] = useState('');
    const [value, setValue] = useState('');
    const [expiration, setExpiration] = useState(5); // Default expiration is 5 seconds

    const handleSet = async (e) => {
        e.preventDefault();
        try {
            await axios.post('/api/cache', { key, value, expiration });
            setKey('');
            setValue('');
            setExpiration(5);
        } catch (error) {
            console.error('Error setting cache:', error);
        }
    };

    return (
        <form onSubmit={handleSet}>
            <div className="cache-input">
                <input
                    type="text"
                    value={key}
                    onChange={(e) => setKey(e.target.value)}
                    placeholder="Key"
                />
                <input
                    type="text"
                    value={value}
                    onChange={(e) => setValue(e.target.value)}
                    placeholder="Value"
                />
                <input
                    type="number"
                    value={expiration}
                    onChange={(e) => setExpiration(Number(e.target.value))}
                    placeholder="Expiration (s)"
                />
                <button type="submit" onClick={handleSet}>Set Cache</button>
            </div>
        </form>
    );
};

export default CacheInput;