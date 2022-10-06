# eFishery Warehouse Management System

# 1. Summary

Warehouse Management System to manage user orders, warehouse inventory, in and out inventory from factory to warehouse and from warehouse to users.

# 2. Problem & Motivation

- Problem

  1. Repeated and tedious bookkeeping

     You could say almost warehouse have to write how much inventory that came to warehouse and out to the users and you need to update it to supervisor in the end of the day. Every item that send to user you need to keep track record of it while writing it down sounds easy for 1 or 10 orders, it would be hard to write manually 100 orders.

  2. Inconsistency through all the warehouse

     Every supervisor and staffs would not have the same consistency while writing down everything even if it’s already referring to some guideline from the main office.

  3. Hard to monitor every warehouse

     If you have many warehouse in Indonesia you can’t check to every warehouse manually going there or just waiting for the report from every supervisor in each warehouse.

  4. Hard to track Inventory

     Every product in inventory need to be organized is it by category, size, weight or even volumes. We need to know where is the item in warehouse.

- Motivation
  By creating such system we can track how much inventory that each warehouse have, when to replenish inventory, when warehouse need to send to customer(user), and for easier sending inventory to user from the nearest warehouse.

# 3. Detailed Design

## Specification

1. Listing Warehouses

   Inventory itu menyimpan ada produk apa, jumlahnya berapa di gudang mana

2. Inventory Management

   Penyimpanan produk didalam warehouse secara proksimitas tata letak posisi.

3. Stock-in/Stock-out

   Stock in stock out itu proses keluar masuk barang ke gudang

## Feature

- Create warehouse user
- Auth warehouse user
- as a supervisor you can add more inventory and got notification if the stock is low
- as a staff you can retrieve item from supplier and send ordered item to users
- as superuser you could see all of this above without restriction
- Ongoing order transaction that is in listed ascending by whos first to order (by default)
- Sort order by quantity, weight, date, order process (accepted, on-process, sent from warehouse)
- as a staff you could save an item into a specific aisle, row, etc of each exact system in every warehouse (every warehouse has a different system to make it easier to access from the most oredered item to least ordered item or from sheer quantity it’s up to the supervisor to make all easy)

## Entity Relationship Diagram

![warehouse_test_efishery - public2.png](eFishery%20Warehouse%20Management%20System%203aab94e0380e4d25852a111205ba3d89/warehouse_test_efishery_-_public2.png)

[efishery.warehouse_roles](https://www.notion.so/15b85e7c369d4a1ca0f1527ed7cc343c)

[efishery.warehouse_categories](https://www.notion.so/25adbbccbeb14214828fa3b676dfb2e5)

[efishery.supplier](https://www.notion.so/a251f565ad6545768e838d4e004081d1)

[efishery.supplier_address](https://www.notion.so/354c4a981c4f48d996af4db3bdde26fb)

[efishery.warehouse_address](https://www.notion.so/4d0e762887454edba687726c10cbffe3)

[efishery.warehouse](https://www.notion.so/7c87375da5cb4386bbd4863857b2e528)

[efishery.warehouse_sections](https://www.notion.so/9d85a8248a3d4d34a4ba4a5ad59b8f8e)

[efishery.warehouse_orders](https://www.notion.so/863406d3c1474497a23289ffa9d0f36a)

[efishery.workers](https://www.notion.so/e0288efea6e843b885064c7a50dbd744)

[efishery.warehouse_products](https://www.notion.so/c4529a540c964272a1154a39064f1054)

[efishery.warehouse_auths](https://www.notion.so/adcb62001d694c26a9c80431e3910338)

## API Contract & Documentation

[API Documentation](https://documenter.getpostman.com/view/16004863/2s83zdwS77)

## Prerequisite

To run this app, you might to ensure youe machine has these instance installed:

1. Go Programming Language (1.18 or higher)
2. Docker with Docker Compose (latest version)
3. Postman or other RestAPI client
4. Code / Text Editor
