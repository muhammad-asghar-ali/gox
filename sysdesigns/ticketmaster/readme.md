# Design a Ticketmaster

Ticketmaster is an online platform that allows users to purchase tickets for concerts, sports events, theater, and other live entertainment.

1.  **Requirements:**

    1. **Functional Requirements:**

       - User should be aboe to view events.
       - User should be able to dearch for events.
       - User should be able to book tickets for events.
       - User should be able to view their booked events.
       - Admin or event coordinators should be able to add events.
       - Popular events should have dynamic pricing.
       - Admin or event coordinators should be able to remove the events.
       - Should should be able to see the avaiable seats for the event.
       - Process payments for book event (optional).
       - Generate and validate seats/tickets (optional).
       - Revers the book event (cancel the booked event) after certain time.

    2. **Non Functional Requirements:**

       - The system should prioritize availability for searching.
       - The system should not book a ticket more the one time for specific user (if the booking event is active).
       - The system should be scalable and able to handle high throughput in the form of popular events (10 million user, one event).
       - The system should have low latency search (< 500ms)
       - The system is read heavy, and thus needs to be able to support high read throughput (100:1).
       - The system should protect user data and adhere to GDPR (General Data Protection Regulation).
       - The system should be fault tolerant.
       - The system should provide secure transactions for purchases.
       - The system should be tested and east to deploy.
       - The system should have regular backups.

2.  **Assumptions:**

3.  **Capacity Estimation:**

    1. **Throughput Requirements:**

    2. **Storage Estimation:**

    3. **Bandwidth Estimation:** (optional)

    4. **Caching Estimation:**

    5. **Infrastructure Sizing:**

4.  **High Level Design:**

5.  **Database Design:**

6.  **System API Design:**

7.  **Deep Dive:**
