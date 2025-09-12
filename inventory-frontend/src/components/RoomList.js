import React, { useState, useEffect } from 'react';
import apiClient from '../utils/apiClient';
import { useParams } from 'react-router-dom';

const RoomList = () => {
    const [rooms, setRooms] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const { id: houseId } = useParams(); // Get house ID from route params

    useEffect(() => {
        fetchRooms();
    }, []);

    const fetchRooms = async () => {
        try {
            let response;
            if (houseId) {
                // Fetch rooms for a specific house
                response = await apiClient.get(`/houses/${houseId}/rooms`);
            } else {
                // Fetch all rooms
                response = await apiClient.get('/rooms');
            }
            // Ensure response.data is an array, default to empty array if null/undefined
            setRooms(response.data || []);
            setLoading(false);
        } catch (err) {
            setError('Failed to fetch rooms');
            setLoading(false);
        }
    };

    if (loading) return <div>Loading rooms...</div>;
    if (error) return <div>{error}</div>;

    return (
        <div>
            <h2>Rooms</h2>
            {houseId && (
                <p><a href={`/houses/${houseId}`}>← Back to House</a></p>
            )}
            {houseId && (
                <p><a href={`/houses/${houseId}/rooms/new`}>Create New Room</a></p>
            )}
            {rooms && rooms.length > 0 ? (
                <table>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Area</th>
                            <th>Unit</th>
                            <th>Created At</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {rooms.map(room => (
                            <tr key={room.id}>
                                <td>{room.name}</td>
                                <td>{room.area || 'N/A'}</td>
                                <td>{room.unit || 'N/A'}</td>
                                <td>{new Date(room.created_at).toLocaleDateString()}</td>
                                <td>
                                    <a href={`/rooms/${room.id}/edit`}>Edit</a>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            ) : (
                <p>No rooms found.</p>
            )}
        </div>
    );
};

export default RoomList;
