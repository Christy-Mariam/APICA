//To display the current cache entries

import React, { useEffect, useState } from 'react';
import axios from 'axios';
import CacheItem from '../CacheItem/CacheItem';


const CacheDisplay = () => {
    const [cache, setCache] = useState([]);

    const fetchCacheItems = async () => {
        try {
            const response = await axios.get('/api/cache');
            setCache(response.data);
        } catch (error) {
            console.error('Error fetching cache:', error);
        }
    }

    useEffect(() => {
        fetchCacheItems();
        const intervalId = setInterval(fetchCacheItems, 5000); // Refresh every 5 seconds
        return () => clearInterval(intervalId);
    }, []);

    return (
        <div className='cache-display'>
            <h2>Current Cache</h2>
            <ul className='cache-list'>
                {cache.map((item) => (
                    <CacheItem 
                    key={item.key} 
                    item={item} 
                    onDelete = {fetchCacheItems}
                    />
                ))}
            </ul>
        </div>
    );
};

export default CacheDisplay;