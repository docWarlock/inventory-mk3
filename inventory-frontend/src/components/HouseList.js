import React, { useState, useEffect } from 'react';
import apiClient from '../utils/apiClient';
import { Link } from 'react-router-dom';

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
            // Handle case where API returns null for empty results
            setHouses(response.data || []);
            setLoading(false);
        } catch (err) {
            setError('Failed to fetch houses');
            setLoading(false);
        }
    };

    const handleDelete = async (houseId) => {
        if (!window.confirm('Are you sure you want to delete this house?')) {
            return;
        }

        try {
            await apiClient.delete(`/houses/${houseId}`);
            // Refresh the list after deletion
            fetchHouses();
        } catch (err) {
            setError('Failed to delete house');
        }
    };

    if (loading) return <div>Loading houses...</div>;
    if (error) return <div>{error}</div>;

    return (
        <div>
            <h2>Houses</h2>
            <p><Link to="/houses/new">Add New House</Link></p>
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
                                <a href={`/houses/${house.id}/rooms`}>View Rooms</a> |
                                <button onClick={() => handleDelete(house.id)} style={{ border: 'none', background: 'none', color: 'red', cursor: 'pointer' }}>
                                    Delete
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default HouseList;
