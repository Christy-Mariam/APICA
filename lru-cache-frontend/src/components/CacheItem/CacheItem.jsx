// A component to display individual cache entries with options to delete them.

import React from 'react';
import axios from 'axios';

const CacheItem = ({ item }) => {
    const handleDelete = async () => {
        try {
            await axios.delete(`/api/cache?key=${item.key}`);
        } catch (error) {
            console.error('Error deleting cache item:', error);
        }
    };

    return (
        <li className='cache-item'>
            <span className='cache-item-key'>{item.key}: </span>
            <span className='cache-item-value'>{item.value} </span>
            <span className='cache-item-expiration'>(Expires in: {item.expiration} s) </span>
            <button onClick={handleDelete}>Delete</button>
        </li>
    );
};

export default CacheItem;