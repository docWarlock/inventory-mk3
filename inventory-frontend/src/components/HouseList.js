import React, { useState, useEffect } from 'react';
import apiClient from '../utils/apiClient';

const HouseList = () => {
    const [houses, setHouses] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        fetchHouses();
    }, []);

    const fetchHouses = async () => {
        try {
            const response = await apiClient.get('/houses');
            setHouses(response.data);
            setLoading(false);
        } catch (err) {
            setError('Failed to fetch houses');
            setLoading(false);
        }
    };

    if (loading) return <div>Loading houses...</div>;
    if (error) return <div>{error}</div>;

    return (
        <div>
            <h2>Houses</h2>
            <table>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Total Area</th>
                        <th>Unit</th>
                        <th>Created At</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {houses.map(house => (
                        <tr key={house.id}>
                            <td>{house.name}</td>
                            <td>{house.total_area || 'N/A'}</td>
                            <td>{house.unit || 'N/A'}</td>
                            <td>{new Date(house.created_at).toLocaleDateString()}</td>
                            <td>
                                <a href={`/houses/${house.id}/edit`}>Edit</a> |
                                <a href={`/houses/${house.id}/rooms`}>View Rooms</a>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default HouseList;
